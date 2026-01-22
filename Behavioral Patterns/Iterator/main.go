package main

import "fmt"

func main() {
	playlist := NewPlaylist()

	playlist.AddSong("Shape of You")
	playlist.AddSong("Bohemian Rhapsody")
	playlist.AddSong("Blinding Lights")

	iterator := playlist.CreateIterator()

	fmt.Println("Now Playing ...")

	for iterator.HasNext() {
		fmt.Println(" 🎵 " + iterator.Next())
	}
}
