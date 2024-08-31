package main

import (
	"fmt"
	"os"

	"github.com/ruslanguns/go-chat/internal/server"
)

func main() {

	port := os.Getenv("PORT")
	server := server.NewServer()

	fmt.Printf("Server is running on :%s...\n", port)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
