package main

import (
	"encoding/json"
	"fmt"
	"game-app/repository/mysql"
	"game-app/service/userservice"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheck)
	mux.HandleFunc("/users/register", userRegisterHandler)
	log.Println("server is listening on port 8080")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Println(server.ListenAndServe())
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"message": "server is running in port 8080"}`))
}

func userRegisterHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(w, `{"message": "invalid method"}`)

		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	var registerReq userservice.RegisterRequest
	err = json.Unmarshal(data, &registerReq)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)
	_, err = userSvc.Register(registerReq)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	w.Write([]byte(`{"message": "user created"}`))
}
