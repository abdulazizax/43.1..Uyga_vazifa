package main

import (
	"server/api"
	"server/storage"
	"server/storage/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := db.DbConnect()
	author := storage.NewUser(db)
	r := api.Router(author)
	fmt.Println("server is listening on 7070")
	err := http.ListenAndServe(":7070", r)
	if err != nil {
		log.Fatal(err)
	}
}
