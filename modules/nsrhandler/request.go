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
	Snssai    string `yaml:"snssai"`
	Ngci      string `yaml:"ngci"`
	Duration  int    `yaml:"duration"`
	Cpu       int    `yaml:"cpu"`
	Bandwidth int    `yaml:"bandwidth"`
}

// ByCpu implements sort.Interface based on the Cpu field.
type ByCpu []SliceList

func (a ByCpu) Len() int           { return len(a) }
func (a ByCpu) Less(i, j int) bool { return a[i].Cpu < a[j].Cpu }
func (a ByCpu) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ByCpu implements sort.Interface based on the Cpu field.
type ByBandwidth []SliceList

func (a ByBandwidth) Len() int           { return len(a) }
func (a ByBandwidth) Less(i, j int) bool { return a[i].Bandwidth < a[j].Bandwidth }
func (a ByBandwidth) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }