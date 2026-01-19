package main

import "fmt"

type CharacterGlyph struct {
	symbol     rune
	fontFamily string
	fontSize   int
	color      string
}

func NewCharacterGlyph(symbol rune, fontFamily string, fontSize int, color string) *CharacterGlyph {
	return &CharacterGlyph{
		symbol:     symbol,
		fontFamily: fontFamily,
		fontSize:   fontSize,
		color:      color,
	}
}

func (cg *CharacterGlyph) Draw(x, y int) {
	fmt.Println("Drawing '" + string(cg.symbol) + "' [Font: " + cg.fontFamily +
		", Size: " + fmt.Sprint(cg.fontSize) + ", Color: " + cg.color + "] at (" + fmt.Sprint(x) + "," + fmt.Sprint(y) + ")")
}
