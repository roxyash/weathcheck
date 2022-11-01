package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"weathcheck/internal/helpers"

	"github.com/Jeffail/gabs"
)

type GeocoderInfo struct {
	Latitude  string `json:"geo_lat"`
	Longitude string `json:"geo_lon"`
	Region    string `json:"region"`
}

//type GraphhopperResponse struct {
//	Hits []Hits `json:"hits"`
//	Locale string `json:"locale"`
//}

//type Hits struct {
//	Point map[string]float64
//	Extend []float64
//	Name
//
//}

func GeocoderDadata(address, token, secret string) ([]GeocoderInfo, error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`[ "%v" ]`, address))
	req, err := http.NewRequest("POST", "https://cleaner.dadata.ru/api/v1/clean/address", data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %v", token))
	req.Header.Set("X-Secret", secret)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = helpers.CheckStatus(resp.StatusCode, bodyText)
	if err != nil {
		return nil, err
	}
	geocoderInfo := []GeocoderInfo{}
	err = json.Unmarshal(bodyText, &geocoderInfo)
	if err != nil {
		return nil, err
	}

	return geocoderInfo, nil
}

func GeocoderGraphhopper(address, apikey string) (GeocoderInfo, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://graphhopper.com/api/1/geocode?q=Пресненская%20набережная&locale=en&key=c077a3ab-c8c9-45c6-9ed6-681d4df0a616", nil)
	if err != nil {
		return GeocoderInfo{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return GeocoderInfo{}, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GeocoderInfo{}, err
	}

	jsonParsed, err := gabs.ParseJSON(bodyText)

	if err != nil {
		return GeocoderInfo{}, err
	}

	gObjLat, err := jsonParsed.JSONPointer("/hits/0/point/lat")
	if err != nil {
		return GeocoderInfo{}, err
	}

	gObjLon, err := jsonParsed.JSONPointer("/hits/0/point/lng")
	if err != nil {
		return GeocoderInfo{}, err
	}

	gObjRegion, err := jsonParsed.JSONPointer("/hits/0/city")
	if err != nil {
		return GeocoderInfo{}, err
	}

	lat, ok := gObjLat.Data().(float64)
	if !ok {
		return GeocoderInfo{}, errors.New("Error unmarshal lat")
	}
	lon, ok := gObjLon.Data().(float64)
	if !ok {
		return GeocoderInfo{}, errors.New("Error unmarshal lat")
	}
	region, ok := gObjRegion.Data().(string)

	if !ok {
		return GeocoderInfo{}, errors.New("Error unmarshal lat")
	}

	geocoderInfo := GeocoderInfo{
		Latitude:  fmt.Sprintf("%v", lat),
		Longitude: fmt.Sprintf("%v", lon),
		Region:    region,
	}

	return geocoderInfo, nil
}
