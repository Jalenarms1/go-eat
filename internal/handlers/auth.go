package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func HandleSignUpPage(w http.ResponseWriter, r *http.Request) error {

	t, err := template.ParseFiles("./internal/templates/auth/signup.html")
	if err != nil {
		return err
	}

	t.Execute(w, nil)

	return nil
}

func HandleNewSignUp(w http.ResponseWriter, r *http.Request) error {

	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(email, password)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("signed-up"))

	return nil

}
