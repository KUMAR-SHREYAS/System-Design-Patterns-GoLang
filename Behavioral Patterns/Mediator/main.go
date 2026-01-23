package main

import "fmt"

func main() {
	var mediator *FormMediator = NewFormMediator()
	usernameField := NewTextField(mediator)
	passwordField := NewTextField(mediator)
	loginButton := NewButton(mediator)
	statusLabel := NewLabel(mediator)

	mediator.setUsernameField(usernameField)
	mediator.setPasswordField(passwordField)
	mediator.setLoginButton(loginButton)
	mediator.setStatusLabel(statusLabel)

	usernameField.setText("admin")
	passwordField.setText("1234")
	loginButton.click()
	fmt.Println("\n ---New Attempt with wrong password---")
	passwordField.setText("wrong")
	loginButton.click()
}
