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

// ByResource implements sort.Interface based on the Resource field.
type ByResource []SliceList

func (a ByResource) Len() int           { return len(a) }
func (a ByResource) Less(i, j int) bool { return a[i].Resource < a[j].Resource }
func (a ByResource) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }