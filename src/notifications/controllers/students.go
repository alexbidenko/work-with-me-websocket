package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pusher/pusher-http-go"
	"net/http"
)

type Authorization struct {
	Login    string
	Password string
}

func InitNotifications(r *mux.Router) {
	r.HandleFunc("/notification/user/{userId}", sendUserMessage).Methods("POST")
	r.HandleFunc("/notification/project/{projectId}", sendProjectMessage).Methods("POST")
}

func sendUserMessage(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]

	title := r.FormValue("title")
	description := r.FormValue("description")

	pusherClient := pusher.Client{
		AppID:   "966947",
		Key:     "8da04f0e1ecfefbeaecc",
		Secret:  "7d92e3ac99cd7e9e6b3f",
		Cluster: "eu",
	}

	message := map[string]string{
		"title": title,
		"description": description,
	}

	err := pusherClient.Trigger("notification-user-" + userId, "messages", message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(message)
	w.Write(response)
}

func sendProjectMessage(w http.ResponseWriter, r *http.Request) {
	projectId := mux.Vars(r)["projectId"]

	title := r.FormValue("title")
	description := r.FormValue("description")

	pusherClient := pusher.Client{
		AppID:   "966947",
		Key:     "8da04f0e1ecfefbeaecc",
		Secret:  "7d92e3ac99cd7e9e6b3f",
		Cluster: "eu",
	}

	message := map[string]string{
		"title": title,
		"description": description,
	}

	err := pusherClient.Trigger("notification-project-" + projectId, "messages", message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(message)
	w.Write(response)
}
