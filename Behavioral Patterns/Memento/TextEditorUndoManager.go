// Caretaker
// The TextEditorUndoManager allows us to perform undo operations cleanly, without the client managing snapshots directly.
package main

import "fmt"

type TextEditorUndoManager struct {
	history []TextEditorMemento
}

func NewTextEditorUndoManager() *TextEditorUndoManager {

	return &TextEditorUndoManager{history: []TextEditorMemento{}}
}

func (t *TextEditorUndoManager) save(editor TextEditor) {
	t.history = append(t.history, *editor.save()) // t.history saves the TextEditorMemento coz it needs to
	// call the getState() method
}
func (t *TextEditorUndoManager) undo(editor TextEditor) {
	sz := len(t.history)
	if sz > 0 {
		pop := t.history[sz-1]
		editor.restore(&pop)
		t.history = t.history[:sz-1]
	} else {
		fmt.Println("Nothing to undo")
	}
}
