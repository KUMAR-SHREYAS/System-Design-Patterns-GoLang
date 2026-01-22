package main

import "fmt"

type SmartButton struct {
	currCommand Command
	history     []Command
}

func NewSmartButton() *SmartButton {
	return &SmartButton{history: []Command{}}
}

func (s *SmartButton) setCommand(command Command) {
	s.currCommand = command
}

func (s *SmartButton) press() {
	if s.currCommand != nil {
		s.currCommand.Execute()
		s.history = append(s.history, s.currCommand)
	} else {
		fmt.Println("No command assigned. ")
	}
}

func (s *SmartButton) undoLast() {
	if len(s.history) > 0 {
		var lastCommand Command = s.history[len(s.history)-1]
		s.history = s.history[:len(s.history)-1]
		lastCommand.Undo()
	} else {
		fmt.Println("Nothing to do")
	}
}
