package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Coords struct {
	Latitude  string `json:"geo_lat"`
	Longitude string `json:"geo_lon"`
	// Source: "москва сухонская 11",
    // "result": "г Москва, ул Сухонская, д 11",
    // "postal_code": "127642",
    // "country": "Россия",
    // "region": "Москва",
    // "city_area": "Северо-восточный",
    // "city_district": "Северное Медведково",
    // "street": "Сухонская",
    // "house": "11",
    // "geo_lat": "55.8782557",
    // "geo_lon": "37.65372",
    // "qc_geo": 0
}

func Geocoder(address, token, secret string) ([]Coords, error) {
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

	coords := []Coords{}
	// var coords interface{}


	err = json.Unmarshal(bodyText, &coords)
	if err != nil {
		return nil, err
	}

	return coords, nil
}
