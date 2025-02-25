package handlers

import (
	"net/http"
	"text/template"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) error {
	subDomain := r.Context().Value(SubDomainCtxKey)

	if subDomain == "dashboard" {
		t, err := template.ParseFiles("./internal/templates/dashboard.html")
		if err != nil {
			return err
		}

		t.Execute(w, nil)

		return nil
	}

	w.Write([]byte(subDomain.(string)))

	return nil
}
