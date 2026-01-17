package main

import (
	"fmt"
	"time"
)

type HighResolutionImage struct {
	fileName  string
	imageData []byte
}

func NewHighResolutionImage(fileName string) *HighResolutionImage {
	HRI := &HighResolutionImage{fileName: fileName}
	HRI.loadImageFromDisk()
	return HRI
}

func (h *HighResolutionImage) loadImageFromDisk() {
	fmt.Println("Loading image: " + h.fileName + " from disk (Expensive Operation) . . .")

	time.Sleep(2 * time.Second)
	h.imageData = make([]byte, 10*1024*1024)
	fmt.Println("Image " + h.fileName + " loaded successfully.")
}

func (h *HighResolutionImage) display() {
	fmt.Println("Displaying image: " + h.fileName)
}

func (h *HighResolutionImage) getFileName() string {
	return h.fileName
}
