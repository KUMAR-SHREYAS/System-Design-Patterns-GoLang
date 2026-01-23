package main

import "fmt"

func main() {
	data := NewReportData()
	csvExporter := NewCsvReportExporter()
	csvExporter.ExportReport(data, "sales_report")

	fmt.Println()

	pdfExporter := NewPdfReportExporter()
	pdfExporter.ExportReport(data, "financial_summary")

}
