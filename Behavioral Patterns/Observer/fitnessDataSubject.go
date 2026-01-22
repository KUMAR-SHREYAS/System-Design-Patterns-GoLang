// subject interface
package main

type fitnessDataSubject interface {
	RegisterObserver(observer FitnessDataObserver)
	RemoveObserver(observer FitnessDataObserver)
	NotifyObservers()
	NewFitnessDataPushed(steps, activeMinutes, calories int)
	dailyReset()
}
