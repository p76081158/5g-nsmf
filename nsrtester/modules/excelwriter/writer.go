package excelwriter

import (
	"os"
	"log"
	"encoding/csv"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func WriteToExcel(dir string, forecast_time string, csv_data [][]string) {
	path := "../slice-requests/" + dir + "/result-" + forecast_time + ".csv"
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