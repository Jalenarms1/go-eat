package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jalenarms1/go-eat/internal/db"
	"github.com/Jalenarms1/go-eat/internal/handlers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}

	if err := db.SetDb(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB Connected")
}

func main() {

	mux := http.NewServeMux()

	registerRouter(mux)

	fmt.Println("We here")

	listenAddr := os.Getenv("LISTEN_ADDR")

	fmt.Printf("http://localhost%s\n", listenAddr)
	http.ListenAndServe(listenAddr, handlers.CtxMiddleWare(mux))
}
