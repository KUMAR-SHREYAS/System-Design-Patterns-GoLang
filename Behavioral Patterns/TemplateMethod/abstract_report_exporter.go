package main

import "fmt"

// Interface for steps that MUST be implemented by child exporters
// These are the methods that change behavior at runtime
type ReportExporter interface {
	WriteHeader(data *ReportData)
	WriteDataRow(data *ReportData)
}

// Abstract base class equivalent
// so if a exporter type(csv ,excel, pdf)
// is of Abstract class then that exporter
// automatically has to implement ReportExporter
// interface
type AbstractReporterExporter struct {
	Exporter ReportExporter
}

//Template method
func (a *AbstractReporterExporter) ExportReport(data *ReportData, filePath string) {
	a.PrepareData(data)
	a.OpenFile(filePath)

	a.Exporter.WriteHeader(data)
	a.Exporter.WriteDataRow(data)

	a.WriteFooter(data)
	a.CloseFile(filePath)

	fmt.Println("Export Complete: ", filePath)
}

func (a *AbstractReporterExporter) PrepareData(data *ReportData) {
	fmt.Println("Preparing report data (common step)...")
}

func (a *AbstractReporterExporter) OpenFile(filePath string) {
	fmt.Println("Opening file '" + filePath + "'")
}

func (a *AbstractReporterExporter) WriteFooter(data *ReportData) {
	fmt.Println("Writing footer (default: no footer).")
}

func (a *AbstractReporterExporter) CloseFile(filePath string) {
	fmt.Println("Closing File '" + filePath + "'")
}
