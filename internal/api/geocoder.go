package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type GeocoderInfo struct {
	Latitude  string `json:"geo_lat"`
	Longitude string `json:"geo_lon"`
	Region    string `json:"region"`
}

func Geocoder(address, token, secret string) ([]GeocoderInfo, error) {
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

	geocoderInfo := []GeocoderInfo{}
	err = json.Unmarshal(bodyText, &geocoderInfo)
	if err != nil {
		return nil, err
	}

	return geocoderInfo, nil
}
