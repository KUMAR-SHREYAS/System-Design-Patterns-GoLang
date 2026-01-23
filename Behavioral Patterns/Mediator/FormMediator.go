package main

type FormMediator struct {
	usernameField *TextField
	passwordField *TextField
	loginButton   *Button
	statusLabel   *Label
}

func NewFormMediator() *FormMediator {
	return &FormMediator{}
}

func (f *FormMediator) setUsernameField(usernameField *TextField) {
	f.usernameField = usernameField
}

func (f *FormMediator) setPasswordField(passwordField *TextField) {
	f.passwordField = passwordField
}
func (f *FormMediator) setLoginButton(loginButton *Button) {
	f.loginButton = loginButton
}
func (f *FormMediator) setStatusLabel(statusLabel *Label) {
	f.statusLabel = statusLabel
}

func (f *FormMediator) ComponentChanged(comp *UIComponent) {
	if comp == &f.usernameField.component || comp == &f.passwordField.component {
		var enableButton bool = len(f.usernameField.getText()) > 0 &&
			len(f.passwordField.getText()) > 0
		f.loginButton.setEnabled(enableButton)
	} else if comp == &f.loginButton.component {
		username := f.usernameField.getText()
		password := f.passwordField.getText()

		if "admin" == username && "1234" == password {
			f.statusLabel.setText("Status: ✅ Login successful!")
			// Optionally, you could disable the button here after success if desired
			f.loginButton.setEnabled(true)
		} else {
			f.statusLabel.setText("Status: ❌ Invalid credentials.")
			// Do NOT disable the button here—keep it enabled for retry
		}
	}
}
