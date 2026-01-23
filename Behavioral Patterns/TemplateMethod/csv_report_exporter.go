package main

import (
	"fmt"
	"strings"
)

type CsvReportExporter struct {
	AbstractReporterExporter
}

func NewCsvReportExporter() *CsvReportExporter {
	csv := &CsvReportExporter{}
	csv.AbstractReporterExporter = AbstractReporterExporter{Exporter: csv}
	return csv
}

func (c *CsvReportExporter) WriteHeader(data *ReportData) {
	fmt.Println("CSV: Writing header:", strings.Join(data.Headers, ","))
}

func (c *CsvReportExporter) WriteDataRow(data *ReportData) {
	fmt.Println("CSV: writing data rows")
	for _, row := range data.Rows {
		values := []any{}
		for _, v := range row {
			values = append(values, v)
		}
		fmt.Println("CSV: ", values)
	}
}
