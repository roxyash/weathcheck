package types

type ResponseWeatherInfo struct {
	Temperature string `json:"temp"`
	Weather     string `json:"weather"`
	Region      string `json:"region"`
	Error       string `json:"error"`
}
