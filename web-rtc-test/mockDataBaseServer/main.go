package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SessionInfo struct {
	SessionId int    `json:"sessionId"`
	Attribute string `json:"attribute"`
}

type UserId struct {
	UserId int `json:"userId"`
}

func main() {
	mockDataOffer := SessionInfo{SessionId: 1, Attribute: "offer"}
	mockDataAnswer := SessionInfo{SessionId: 1, Attribute: "answer"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed.", http.StatusMethodNotAllowed)
			return
		}

		var userId UserId
		err := json.NewDecoder(r.Body).Decode(&userId)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var response SessionInfo
		if userId.UserId%2 == 1 {
			response = mockDataOffer
		} else {
			response = mockDataAnswer
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	port := "3002"
	fmt.Printf("server: http://localhost:%s/\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}
}
