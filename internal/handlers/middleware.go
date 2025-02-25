package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type CtxKey string

var SubDomainCtxKey CtxKey = "target_subDomain"
var OwnerSessionKey CtxKey = "owner_session_myapp"

func CtxMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Man in the middle")

		subDomain := strings.Split(r.Host, ".")[0]

		cookie, err := r.Cookie(string(OwnerSessionKey))

		if err != nil {
			fmt.Println("No cookie found")
		}

		fmt.Println(subDomain)

		ctx := context.WithValue(r.Context(), SubDomainCtxKey, subDomain)
		ctx = context.WithValue(ctx, OwnerSessionKey, cookie)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
