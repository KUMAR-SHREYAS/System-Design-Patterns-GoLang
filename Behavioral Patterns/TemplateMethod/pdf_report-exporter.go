package main

import (
	"fmt"
	"strings"
)

type PdfReportExporter struct {
	AbstractReporterExporter
}

func NewPdfReportExporter() *PdfReportExporter {
	Pdf := &PdfReportExporter{}
	Pdf.AbstractReporterExporter = AbstractReporterExporter{Exporter: Pdf}
	return Pdf
}

func (c *PdfReportExporter) WriteHeader(data *ReportData) {
	fmt.Println("PDF: Writing header:", strings.Join(data.Headers, ","))
}

func (c *PdfReportExporter) WriteDataRow(data *ReportData) {
	fmt.Println("PDF: writing data rows")
	for _, row := range data.Rows {
		values := []any{}
		for _, v := range row {
			values = append(values, v)
		}
		fmt.Println("PDF: ", values)
	}
}
