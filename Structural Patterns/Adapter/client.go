package main

import "fmt"

type Client struct{}

func (c *Client) InsertLightningPortIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
