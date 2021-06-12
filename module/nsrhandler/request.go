package nsrhandler

// Yaml2Go
type Yaml2GoRequestList struct {
	RequestList RequestList `yaml:"requestList"`
}

// RequestList
type RequestList struct {
	SliceList []SliceList `yaml:"sliceList"`
}

// SliceList
type SliceList struct {
	Snssai   string `yaml:"snssai"`
	Ngci     string `yaml:"ngci"`
	Duration int    `yaml:"duration"`
	Resource int    `yaml:"resource"`
}