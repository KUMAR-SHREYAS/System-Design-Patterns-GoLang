package main

import "fmt"

type GoalNotifier struct {
	stepGoal    int
	goalReached bool
}

func NewGoalNotifier() *GoalNotifier {
	return &GoalNotifier{stepGoal: 10000}
}

func (g *GoalNotifier) Update(data FitnessData) {
	if data.GetSteps() >= g.stepGoal && !g.goalReached {
		fmt.Printf("Notifier → 🎉 Goal Reached! You've hit %d steps!", g.stepGoal)
		g.goalReached = true
	}
}

func (g *GoalNotifier) reset() {
	g.goalReached = false
}
