package main

import (
	"fmt"
	"strconv"
)

type LiveActivityDisplay struct {
}

func NewLiveActivityDisplay() *LiveActivityDisplay {
	return &LiveActivityDisplay{}
}

func (l *LiveActivityDisplay) Update(data FitnessData) {
	fmt.Println("Live Display → Steps: " + strconv.Itoa(data.GetSteps()) +
		" | Active Minutes: " + strconv.Itoa(data.GetActiveMinutes()) +
		" | Calories: " + strconv.Itoa(data.GetCalories()))
}
