package main

import "fmt"

func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "musket":
		return newMusket(), nil
	case "ak47":
		return newAK47(), nil
	default:
		return nil, fmt.Errorf("unknown gun type: %s", gunType)
	}
}
