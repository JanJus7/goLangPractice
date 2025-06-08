package tools

type Config struct {
    HighTemperatureThreshold float64 `json:"highTemperatureThreshold"`
	LowTemperatureThreshold  float64 `json:"lowTemperatureThreshold"`
    WindThreshold        float64 `json:"windThreshold"`
    RainThreshold        float64 `json:"rainThreshold"`
}