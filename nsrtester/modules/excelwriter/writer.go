package excelwriter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func WriteToExcel(dir string, resource string, sort bool, forecast_time string, concat bool, csv_data [][]string) {
	path := "../slice-requests/" + dir + "/result-" + resource + "-sort-" + strconv.FormatBool(sort) + "-forecast-"+ forecast_time + "-concat-" + strconv.FormatBool(concat) + ".csv"
	file, err := os.Create(path)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range csv_data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}