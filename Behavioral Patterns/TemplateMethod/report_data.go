package main

type ReportData struct {
	Headers []string
	Rows    []map[string]any
}

func NewReportData() *ReportData {
	return &ReportData{
		Headers: []string{"ID", "Name", "Amount"},
		Rows: []map[string]any{
			{"ID": 1, "Name": "Alice", "Amount": 1200},
			{"ID": 2, "Name": "Bob", "Amount": 3400},
		},
	}
}
