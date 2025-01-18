package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var serverStore = make([]Server, 3)

type SNSRequest struct {
	ServerID int64 `json:"server"`
	Order int64 `json:"order"`
}

type Message struct {
	ServerID int64 `json:"id"`
	Order int64 `json:"order"`
}

type Server struct {
	ServerID int64 `json:"id"`
	ServerURL string `json:"url"`
}

type ServerStore []Server

func main() {
	mux := http.NewServeMux()
	mux.Handle("/send", sendFunc())
	mux.Handle("/register/server", registerServerFunc())
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func sendFunc() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m, err := packRequestToMessage(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		url, err := getTargetServer(m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := sendToServer(m, url); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("%#v", m)
	})
}

func registerServerFunc() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// サーバーを登録する
		if err := registerServer(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})
}

func packRequestToMessage(r *http.Request) (*Message, error) {
	var sr SNSRequest
	if err := json.NewDecoder(r.Body).Decode(&sr); err != nil {
		return nil, err
	}
	return &Message{
		ServerID: sr.ServerID,
		Order: sr.Order,
	}, nil
}

func registerServer(r *http.Request) error {
	var s Server
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		return err
	}
	serverStore = append(serverStore, s)
	return nil
}

// 対象サーバーにメッセージを送信する
func sendToServer(m *Message, url string) error {
	message, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error occur", err)
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("%#v",req)
	fmt.Println("status code:", resp.StatusCode)
	return nil
}

func getTargetServer(m *Message) (string, error) {
	for _, v := range serverStore {
		if v.ServerID == m.ServerID {
			return v.ServerURL, nil
		}
	}
	return "", fmt.Errorf("there is no target server")
}