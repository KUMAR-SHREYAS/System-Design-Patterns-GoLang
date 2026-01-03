package data

// Maps from data/Data.java and global/EventEnum.java
// Define Data struct and Event constants here.

type Data struct {
	ID      string      `json:"id"`
	Payload interface{} `json:"payload"`
	Source  string      `json:"source"`
}
