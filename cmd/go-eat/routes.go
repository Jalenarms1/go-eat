package main

import (
	"fmt"
	"net/http"

	"github.com/Jalenarms1/go-eat/internal/handlers"
)

func registerRouter(mux *http.ServeMux) {

	mux.HandleFunc("/sign-up", catchErrorHandlerFunc(handlers.HandleSignUpPage))

	mux.Handle("/create-food-shop", dashboardRoute(catchErrorHandlerFunc(handlers.HandleNewShop)))

}

func dashboardRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		subDomain := r.Context().Value(handlers.SubDomainCtxKey).(string)

		if subDomain != "dashboard" {
			http.Error(w, "path not accesible from outside of the dashboard space", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type HttpErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func catchErrorHandlerFunc(fn HttpErrorHandlerFunc) http.HandlerFunc {
	fmt.Println("catch")
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}

// func authRouteHandlerFunc(fn HttpErrorHandlerFunc, method string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != method {
// 			fmt.Println("method not allowed")
// 			http.Error(w, "method not allowed", http.StatusBadRequest)
// 			return
// 		}

// 		cookie := r.Context().Value(handlers.OwnerSessionKey).(*http.Cookie)
// 		subDomain := r.Context().Value(handlers.SubDomainCtxKey).(string)

// 		if cookie == nil && subDomain == "dashboard" {
// 			http.Redirect(w, r, "/sign-up", http.StatusPermanentRedirect)
// 			return
// 		}

// 		fmt.Printf("Cookie: %s\n", r.Context().Value(handlers.OwnerSessionKey))
// 		fmt.Println(r.Context().Value(handlers.SubDomainCtxKey))

// 		if err := fn(w, r); err != nil {
// 			fmt.Println(err)
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 	}
// }
