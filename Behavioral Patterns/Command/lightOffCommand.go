package main

type LightOffCommand struct {
	light Light
}

func NewLightOffCommand(light Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (L *LightOffCommand) Execute() {
	L.light.off()
}

func (L *LightOffCommand) Undo() {
	L.light.on()
}
