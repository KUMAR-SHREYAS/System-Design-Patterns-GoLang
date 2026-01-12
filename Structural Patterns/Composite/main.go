package main

import "fmt"

func main() {
	file1 := NewFile("readme.txt", 5)
	file2 := NewFile("photo.jpeg", 1500)
	file3 := NewFile("data.csv", 300)

	documents := NewFolder("Documents")
	documents.addChild(file1)
	documents.addChild(file3)

	pictures := NewFolder("Pictures")
	pictures.addChild(file2)

	home := NewFolder("Home")
	home.addChild(documents)
	home.addChild(pictures)

	fmt.Println("-----File Structure-----")
	home.printStructure("")

	fmt.Printf("\nTotal Size: %d KB", home.getSize())

	fmt.Println("\n---- Deleting All ----")
	home.delete()
}
