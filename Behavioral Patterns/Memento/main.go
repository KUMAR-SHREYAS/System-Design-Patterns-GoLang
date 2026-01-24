package main

import "fmt"

func main() {
	editor := NewTextEditor() // create the originator
	// this will create text and type it. it has save() and restore()
	// methods defined but it will take help of a manager to
	// execute those methods
	manager := NewTextEditorUndoManager()

	editor.Type("Hello")
	manager.save(*editor) //save state : Hello

	editor.Type(" World")
	manager.save(*editor) //save state : Hello World

	editor.Type("!")
	// manager.save(*editor) //save state : Hello World !
	fmt.Println("Current Content: " + editor.getContent())

	fmt.Println("\n--- Undo1 ---")
	manager.undo(*editor) // Back to: Hello World

	fmt.Println("\n--- Undo2 ---")
	manager.undo(*editor) // Back to: Hello

	fmt.Println("\n--- Undo3 ---")
	manager.undo(*editor) // Nothing left to undo
}
