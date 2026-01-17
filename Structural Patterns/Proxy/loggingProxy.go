package main

import (
	"fmt"
	"time"
)

type LoggingProxy struct {
	fileName  string
	realImage *HighResolutionImage
}

func NewLoggingProxy(fileName string) *LoggingProxy {
	LP := &LoggingProxy{fileName: fileName}
	fmt.Println("LoggingProxy: Created for " + LP.fileName + ". Real image not loaded yet.")
	return LP
}
func (LP *LoggingProxy) getFileName() string {
	return LP.fileName
}

func (LP *LoggingProxy) display() {
	fmt.Println("LoggingProxy: Attempting to display " + LP.fileName + " at " + time.Now().Format("2006-01-02 15:04:05"))
	if LP.realImage == nil {
		fmt.Println("LoggingProxy: display() requested for " + LP.fileName + ". Loading high-resolution image...")

		LP.realImage = NewHighResolutionImage(LP.fileName)
	} else {
		fmt.Println("LoggingProxy: Using cached high-resolution image for " + LP.fileName)
	}
	LP.realImage.display()
	fmt.Println("LoggingProxy: Finished displaying " + LP.fileName + " at " + time.Now().Format("2006-01-02 15:04:05"))
}
