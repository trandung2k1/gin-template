package main

import (
	"fmt"
	"gin-gonic/internal/server"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	server := server.NewServer()
	s := fmt.Sprintf("Server listen on:http://localhost:%s", port)
	fmt.Println(s)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
