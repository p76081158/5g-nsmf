package yamltopara

// Yaml2Go
type Yaml2Go struct {
	DatasetInfo DatasetInfo `yaml:"datasetInfo"`
}

// DatasetInfo
type DatasetInfo struct {
	Name       string   `yaml:"name"`
	NgciList   []string `yaml:"ngciList"`
	SliceNum   int      `yaml:"sliceNum"`
	Resource   Resource `yaml:"resource"`
	Timewindow int      `yaml:"timewindow"`
}

// Resource
type Resource struct {
	Name   string `yaml:"name"`
	Limit  int    `yaml:"limit"`
	Lambda int    `yaml:"lambda"`
}