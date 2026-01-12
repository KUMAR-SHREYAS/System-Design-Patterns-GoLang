package main

type FileSystemInterface interface {
	getSize() int
	printStructure(indent string)
	delete()
}
