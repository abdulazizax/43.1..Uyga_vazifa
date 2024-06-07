package api

import (
	"log"
	"net/http"
	"net/rpc"
	"server/api/handler"
	"server/service"
	"server/storage"
)

func Router(st *storage.User) http.Handler {
	client, err := rpc.DialHTTP("tcp", ":5555")
	if err != nil {
		log.Println(err)
	}
	newUser := service.NewUserService(client)
	mux := http.NewServeMux()
	handler := handler.NewHandler(st, newUser)

	mux.HandleFunc("POST /user", handler.CreateUser)
	mux.HandleFunc("GET /user/id", handler.GetUserByID)
	mux.HandleFunc("GET /users", handler.GetAllUsers)
	mux.HandleFunc("PUT /user/id", handler.UpdateUserByID)
	mux.HandleFunc("DELETE /user/id", handler.DeleteUserByID)

	return mux
}
