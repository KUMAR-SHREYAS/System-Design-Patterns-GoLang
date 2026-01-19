package main

import "fmt"

type CharacterFlyweightFactory struct {
	flyweightMap map[string]CharacterFlyweight
}

func NewCharacterFlyweightFactory() *CharacterFlyweightFactory {
	return &CharacterFlyweightFactory{
		flyweightMap: make(map[string]CharacterFlyweight),
	}
}

func (f *CharacterFlyweightFactory) GetFlyweightCount() int {
	return len(f.flyweightMap)
}

func (f *CharacterFlyweightFactory) GetFlyweight(symbol rune, fontFamily string, fontSize int, color string) CharacterFlyweight {
	key := fmt.Sprintf("%c%s%d%s", symbol, fontFamily, fontSize, color)

	if _, ok := f.flyweightMap[key]; !ok {
		f.flyweightMap[key] = &CharacterGlyph{
			symbol:     symbol,
			fontFamily: fontFamily,
			fontSize:   fontSize,
			color:      color,
		}
	}
	return f.flyweightMap[key]
}
