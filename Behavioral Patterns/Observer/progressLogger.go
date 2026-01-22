package main

import (
	"fmt"
	"strconv"
)

type ProgressLogger struct {
}

func NewProgressLogger() *ProgressLogger {
	return &ProgressLogger{}
}

func (pl *ProgressLogger) Update(data FitnessData) {
	fmt.Println("Logger → Saving to DB: Steps= " + strconv.Itoa(data.GetSteps()) +
		" | Active Minutes: " + strconv.Itoa(data.GetActiveMinutes()) +
		" | Calories: " + strconv.Itoa(data.GetCalories()))
}
