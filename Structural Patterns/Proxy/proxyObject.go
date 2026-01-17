package main

import "fmt"

type ImageProxy struct {
	fileName  string
	realImage *HighResolutionImage
}

func NewImageProxy(fileName string) *ImageProxy {
	IP := &ImageProxy{fileName: fileName}
	fmt.Println("ImageProxy: Created for " + IP.fileName + ". Real image not loaded yet.")
	return IP
}
func (IP *ImageProxy) getFileName() string {
	return IP.fileName
}

func (IP *ImageProxy) display() {
	if IP.realImage == nil {
		fmt.Println("ImageProxy: display() requested for " + IP.fileName + ". Loading high-resolution image...")

		IP.realImage = NewHighResolutionImage(IP.fileName)
	} else {
		fmt.Println("ImageProxy: Using cached high-resolution image for " + IP.fileName)
	}
	IP.realImage.display()
}
