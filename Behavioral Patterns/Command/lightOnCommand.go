package main

type LightOnCommand struct {
	command Command
	light   Light
}

func NewLightOnCommand(light Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (L *LightOnCommand) Execute() {
	L.light.on()
}

func (L *LightOnCommand) Undo() {
	L.light.off()
}
