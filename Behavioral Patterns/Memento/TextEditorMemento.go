// Memento Class
// This class is passive and doesn’t contain any logic — just a snapshot of the editor’s state.
package main

type TextEditorMemento struct {
	state string
}

func NewTextEditorMemento(state string) *TextEditorMemento {
	return &TextEditorMemento{state: state}
}

func (t *TextEditorMemento) getState() string {
	return t.state
}
