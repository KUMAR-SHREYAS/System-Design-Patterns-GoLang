package main

import "fmt"

func main() {
	fmt.Println("Application Started. Initializing image proxies for gallery...")
	// Create lightweight proxies instead of full image objects

	// image1 := NewImageProxy("photo1.jpg") this directlyb exposes image proxy to the clien main,go
	// which is not required , as client should not know if interacting with proxy or real object
	// also this exposes proxy's internal structure to the client, not required

	var image1 Image = NewImageProxy("photo1.jpg")
	// image2 := NewImageProxy("photo2.png")
	var image3 Image = NewImageProxy("photo3.gif")

	fmt.Println("Gallery initialized. No images actually loaded yet.")
	fmt.Println("Image 1 File Name: " + image1.getFileName()) //does not trigger actual image, just a proxy
	fmt.Println()
	fmt.Println()

	//User clicks on image 1, and requests to display image 1
	fmt.Println("\nUser requests to display: " + image1.getFileName())
	image1.display() // Lazy loading, first time realImage object created
	fmt.Println()
	fmt.Println()

	//Again user clicks on image 1, this time cached content is diplayed
	fmt.Println("\nUser requests to display: " + image1.getFileName())
	image1.display()
	fmt.Println()
	fmt.Println()

	fmt.Println("\nUser requests to display: " + image3.getFileName())
	image3.display()

	fmt.Println("\nApplication finished. Note: photo2.png was never loaded.")

}
