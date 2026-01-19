package main

func main() {
	editor := NewTextEditorClient()
	word := "Hello"

	for i, char := range word {
		editor.AddCharacter(char, i*15+10, 50, "Arial", 14, "#000000")
	}
	word2 := "World"
	for i, char := range word2 {
		editor.AddCharacter(char, i*15+10, 100, "Times New Roman", 14, "#3333FF")
	}
	editor.RenderDocument()
}
