package nsrhandler

// Yaml2Go
type Yaml2GoForecastingBlock struct {
	ForecastingBlock []ForecastingBlock `yaml:"forecastingBlock"`
}

// // ForecastingBlock
// type ForecastingBlock struct {
// 	Block    int `yaml:"block"`
// 	Duration int `yaml:"duration"`
// 	Resource int `yaml:"resource"`
// }

// ForecastingBlock
type ForecastingBlock struct {
	Block     int `yaml:"block"`
	Duration  int `yaml:"duration"`
	Cpu       int `yaml:"cpu"`
	Bandwidth int `yaml:"bandwidth"`
}