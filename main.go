package main

import (
	"awesomeProject/login-register-demo/helpers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""

	myRouter := mux.NewRouter()

	// signup
	myRouter.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		// Empty data checking
		uNameCheck := helpers.IsEmpty(uName)
		emailCheck := helpers.IsEmpty(email)
		pwdCheck := helpers.IsEmpty(pwd)
		pwdConfirmCheck := helpers.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty string")
			return
		}

		if pwd == pwdConfirm {
			// Save to database (username, email and password)
			fmt.Fprintln(w, "Registration successful.")
		} else {
			fmt.Fprintln(w, "Password information must be the same.")
		}
	})

	// login
	myRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")

		// Empty data checking
		emailCheck := helpers.IsEmpty(email)
		pwdCheck := helpers.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty string")
			return
		}

		dbEmail := "azhar@gmail.com"	// Database simulation
		dbPwd := "12345"				// Database simulation

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(w, "Login successful!")
		} else {
			fmt.Fprintln(w, "Login failed!")
		}
	})

	http.ListenAndServe(":8010", myRouter)
}
