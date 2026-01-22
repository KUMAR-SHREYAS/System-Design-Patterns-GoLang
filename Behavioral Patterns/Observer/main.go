package main

func main() {
	var fitnessData fitnessDataSubject = NewFitnessDataClass()

	display := NewLiveActivityDisplay()
	logger := NewProgressLogger()
	notifier := NewGoalNotifier()

	// Register observers
	fitnessData.RegisterObserver(display)
	fitnessData.RegisterObserver(logger)
	fitnessData.RegisterObserver(notifier)

	// Simulate updates
	fitnessData.NewFitnessDataPushed(500, 5, 20)
	fitnessData.NewFitnessDataPushed(9800, 85, 350)
	fitnessData.NewFitnessDataPushed(10100, 90, 380) // Goal should trigger

	notifier.reset()
	fitnessData.dailyReset()
}
