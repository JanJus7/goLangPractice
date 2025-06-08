package tools

type Config struct {
    TemperatureThreshold float64 `json:"temperatureThreshold"`
    WindThreshold        float64 `json:"windThreshold"`
    RainThreshold        float64 `json:"rainThreshold"`
}