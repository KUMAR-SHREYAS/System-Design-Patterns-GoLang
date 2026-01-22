package main

import "fmt"

type FitnessData struct {
	steps         int
	activeMinutes int
	calories      int
	observers     []FitnessDataObserver
	// subject       fitnessDataSubject
}

func NewFitnessDataClass() *FitnessData {
	return &FitnessData{
		steps:         0,
		activeMinutes: 0,
		calories:      0,
		observers:     []FitnessDataObserver{},
	}
}

func (f *FitnessData) RegisterObserver(observer FitnessDataObserver) {
	f.observers = append(f.observers, observer)
}

func (f *FitnessData) RemoveObserver(observer FitnessDataObserver) {
	for i, obs := range f.observers {
		if obs == observer {
			f.observers = append(f.observers[:i], f.observers[i+1:]...)
			break
		}
	}
}

func (f *FitnessData) NotifyObservers() {
	for _, observer := range f.observers {
		observer.Update(*f)
	}
}

func (f *FitnessData) NewFitnessDataPushed(steps, activeMinutes, calories int) {
	f.calories = calories
	f.steps = steps
	f.activeMinutes = activeMinutes
	fmt.Printf("\nFitnessData: New data received – Steps: %d, Active Minutes: %d, Calories: %d\n", steps, activeMinutes, calories)
	f.NotifyObservers()
}

func (f *FitnessData) dailyReset() {
	f.calories = 0
	f.steps = 0
	f.activeMinutes = 0
	fmt.Println("\nFitnessData: Daily reset performed.")
	f.NotifyObservers()
}
func (f *FitnessData) GetSteps() int {
	return f.steps
}

func (f *FitnessData) GetActiveMinutes() int {
	return f.activeMinutes
}
func (f *FitnessData) GetCalories() int {
	return f.calories
}
