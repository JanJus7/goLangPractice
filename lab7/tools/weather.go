package tools

type Weather struct {
    Daily struct {
        Time             []string  `json:"time"`
        TemperatureMax   []float64 `json:"temperature_2m_max"`
        TemperatureMin   []float64 `json:"temperature_2m_min"`
        PrecipitationSum []float64 `json:"precipitation_sum"`
        WindSpeedMax     []float64 `json:"wind_speed_10m_max"`
    } `json:"daily"`
}

