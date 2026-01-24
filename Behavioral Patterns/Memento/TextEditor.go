// Originator
package main

import "fmt"

type TextEditor struct {
	content string
}

func NewTextEditor() *TextEditor {
	return &TextEditor{}
}

func (t *TextEditor) Type(newText string) {
	t.content += newText
}

func (t *TextEditor) getContent() string {
	return t.content
}

func (t *TextEditor) save() *TextEditorMemento {
	fmt.Println("Saving State: \"" + t.content + "\"")
	return NewTextEditorMemento(t.content)
}
func (t *TextEditor) restore(memento *TextEditorMemento) {
	content := memento.getState()
	fmt.Println("Restore state to: \"" + content + "\"")
}
