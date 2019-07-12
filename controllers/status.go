package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/context"

	"golang-backend/consts"
)

type status struct {
	Name            string `json:"name,omitempty"`
	Version         string `json:"version,omitempty"`
	Uptime          string `json:"uptime,omitempty"`
	EnvironmentName string `json:"environmentName,omitempty"`
}

func Status(w http.ResponseWriter, r *http.Request) {
	uptime := context.Get(r, consts.UptimeKey)

	status := &status{
		os.Getenv("APPLICATION_NAME"),
		os.Getenv("APPLICATION_VERSION"),
		fmt.Sprint(uptime),
		os.Getenv("ENVIRONMENT_NAME"),
	}
	obj, err := json.Marshal(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": 500}\n"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(obj)
	}
}
