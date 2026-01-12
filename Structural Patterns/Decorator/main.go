package main

import "fmt"

func main() {
	text := NewPlainTextView("Hello world")

	fmt.Println("Plain: ")
	text.Render()
	fmt.Println()

	fmt.Println("Bold: ")
	boldText := NewBoldDecorator(text)
	boldText.Render()
	fmt.Println()

	fmt.Println("Italic + Underline")
	italicUnderlineText := NewUnderlineDecorator(
		NewItalicDecorator(text),
	)
	italicUnderlineText.Render()
	fmt.Println()

	fmt.Println("Underline + Italic")
	underlineItalicText := NewItalicDecorator(NewUnderlineDecorator(text))
	underlineItalicText.Render()
	fmt.Println()
	fmt.Println("Bold  + Italic + Underline")
	allStyles := NewUnderlineDecorator(
		NewItalicDecorator(
			NewBoldDecorator(text),
		),
	)
	allStyles.Render()
	fmt.Println()
}
