package yamltopara

// Get dataset info
func GetDataSetInfo(dir string) DatasetInfo {
    var datasetinfo Yaml2GoRequestList
    path := "slice-requests/" + dir + "/DataSet-parameter.yaml"
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &datasetinfo)
    if err != nil {
        panic(err)
    }

    return datasetinfo.DatasetInfo
}