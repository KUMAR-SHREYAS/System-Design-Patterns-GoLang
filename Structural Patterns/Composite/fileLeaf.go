package main

import "fmt"

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) getSize() int {
	return f.size
}

func (f *File) printStructure(indent string) {
	fmt.Println(indent + "-" + f.name + " (" + fmt.Sprintf("%d", f.size) + " KB)")
}

func (f *File) delete() {
	fmt.Println("Deleting file: " + f.name)
}
