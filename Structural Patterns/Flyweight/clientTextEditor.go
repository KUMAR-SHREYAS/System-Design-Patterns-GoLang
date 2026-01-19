package main

import "fmt"

type TextEditorClient struct {
	factory  *CharacterFlyweightFactory
	document []RenderedCharacter
}

func NewTextEditorClient() *TextEditorClient {
	return &TextEditorClient{
		factory:  NewCharacterFlyweightFactory(),
		document: make([]RenderedCharacter, 0),
	}
}

func (c *TextEditorClient) AddCharacter(char rune, x, y int, font string, size int, color string) {
	glyph := c.factory.GetFlyweight(char, font, size, color)
	c.document = append(c.document, RenderedCharacter{
		glyph: glyph,
		x:     x,
		y:     y,
	})
}

func (c *TextEditorClient) RenderDocument() {
	for _, rc := range c.document {
		rc.Render()
	}
	fmt.Println("Total flyweight objects used:", c.factory.GetFlyweightCount())
}

type RenderedCharacter struct {
	glyph CharacterFlyweight
	x     int
	y     int
}

func (rc *RenderedCharacter) Render() {
	rc.glyph.Draw(rc.x, rc.y)
}
