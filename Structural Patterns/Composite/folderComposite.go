package main

import "fmt"

type Folder struct {
	name     string
	children []FileSystemInterface
}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}

func (f *Folder) addChild(child FileSystemInterface) {
	f.children = append(f.children, child)
}

func (f *Folder) getSize() int {
	totalSize := 0
	for _, children := range f.children {
		totalSize += children.getSize()
	}
	return totalSize
}

func (f *Folder) printStructure(indent string) {
	fmt.Println(indent + "+ " + f.name + "/")
	for _, children := range f.children {
		children.printStructure(indent + "+ ")
	}
}

func (f *Folder) delete() {
	for _, children := range f.children {
		children.delete()
	}
	fmt.Println("Deleting folder: " + f.name)
}
