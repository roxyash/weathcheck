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

func getRegion(jsonParsed *gabs.Container) (string, error) {

	gObjRegion, err := jsonParsed.JSONPointer("/hits/0/city")
	if err != nil {
		gObjState, err := jsonParsed.JSONPointer("/hits/0/state")
		if err != nil {
			gObjName, err := jsonParsed.JSONPointer("/hits/0/name")
			if err != nil {
				return "", errors.New("Can`t unmarshal json")
			} else {
				region, ok := gObjName.Data().(string)
				if !ok {
					return "", errors.New("Can`t unmarshal json")
				} else {
					return region, nil
				}
			}
		} else {
			region, ok := gObjState.Data().(string)
			if !ok {
				return "", errors.New("Can`t unmarshal json")
			} else {
				return region, nil
			}
		}
	} else {
		region, ok := gObjRegion.Data().(string)
		if !ok {
			return "", errors.New("Can`t unmarshal json")
		} else {
			return region, nil
		}
	}
}

func GeocoderGraphhopper(address, apikey string) (GeocoderInfo, error) {
	var newAddress string
	checkAddr := strings.Split(address, " ")
	if len(checkAddr) == 0 {
		newAddress = checkAddr[0]
	} else {
		newAddress = strings.Join(checkAddr[:], "%20")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://graphhopper.com/api/1/geocode?q=%v&locale=en&key=%v", newAddress, apikey), nil)
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

	fmt.Printf("bodyText: %v", string(bodyText[:]))

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

	lat, ok := gObjLat.Data().(float64)
	if !ok {
		return GeocoderInfo{}, errors.New("Error unmarshal lat")
	}
	lon, ok := gObjLon.Data().(float64)
	if !ok {
		return GeocoderInfo{}, errors.New("Error unmarshal lon")
	}

	region, err := getRegion(jsonParsed)
	if err != nil {
		return GeocoderInfo{}, errors.New("Error unmarshal region with getRegion")
	}

	geocoderInfo := GeocoderInfo{
		Latitude:  fmt.Sprintf("%v", lat),
		Longitude: fmt.Sprintf("%v", lon),
		Region:    region,
	}

	return geocoderInfo, nil
}
