package yamltopara

// Yaml2Go
type Yaml2GoDatasetInfo struct {
	DatasetInfo DatasetInfo `yaml:"datasetInfo"`
}

// DatasetInfo
type DatasetInfo struct {
	Name              string   `yaml:"name"`
	NgciList          []string `yaml:"ngciList"`
	SliceNum          int      `yaml:"sliceNum"`
	ExtraRequest      int      `yaml:"extraRequest"`
	TestNum           int      `yaml:"testNum"`
	Resource          Resource `yaml:"resource"`
	Timewindow        int      `yaml:"timewindow"`
	ForecastBlockSize int      `yaml:"forecastBlockSize"`
}

// Resource
type Resource struct {
	Cpu       Cpu       `yaml:"cpu"`
	Bandwidth Bandwidth `yaml:"bandwidth"`
	Duration  int       `yaml:"duration"`
}

// Cpu
type Cpu struct {
	Limit    int     `yaml:"limit"`
	Lambda   int     `yaml:"lambda"`
	Discount float64 `yaml:"discount"`
}

// Bandwidth
type Bandwidth struct {
	Limit    int     `yaml:"limit"`
	Lambda   int     `yaml:"lambda"`
	Discount float64 `yaml:"discount"`
}