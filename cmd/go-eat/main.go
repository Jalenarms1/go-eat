package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jalenarms1/go-eat/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}

	mux := http.NewServeMux()

	registerRouter(mux)

	fmt.Println("We here")

	listenAddr := os.Getenv("LISTEN_ADDR")

	fmt.Printf("http://localhost%s\n", listenAddr)
	http.ListenAndServe(listenAddr, handlers.CtxMiddleWare(mux))
}
