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
	Target            string   `yaml:"target"`
	Sort              bool     `yaml:"sort"`
	Concat            bool     `yaml:"concat"`
	Timewindow        int      `yaml:"timewindow"`
	ForecastingTime   int      `yaml:"forecastingTime"`
	ForecastBlockSize int      `yaml:"forecastBlockSize"`
	Regenerate        bool     `yaml:"regenerate"`
}

// Resource
type Resource struct {
	Cpu       Cpu       `yaml:"cpu"`
	Bandwidth Bandwidth `yaml:"bandwidth"`
	Duration  int       `yaml:"duration"`
	Random    bool      `yaml:"random"`
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