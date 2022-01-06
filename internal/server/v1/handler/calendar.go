package handler

import (
	"calendar/internal/storage/db"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var Database *db.Database

func CreateEvent(rw http.ResponseWriter, r *http.Request) {
	var event db.EventDB
	err := json.NewDecoder(r.Body).Decode(&event)
	logrus.Info(event)
	event, err = Database.CreateEvent(event.Title, event.Date)
	if err != nil {
		rw.WriteHeader(500)
		logrus.Error("DB error :" + err.Error())
		return
	}

	event.Date = strings.Replace(event.Date, "T", " ", 10)
	response, err := json.Marshal(event)
	if err != nil {
		rw.WriteHeader(500)
		logrus.Error("Json marshall error :" + err.Error())
		return
	}

	rw.WriteHeader(200)
	_, err = rw.Write(response)
	if err != nil {
		logrus.Error("error :" + err.Error())
		return
	}
}

func UpdateEvent(rw http.ResponseWriter, r *http.Request) {
	var event db.EventDB
	err := json.NewDecoder(r.Body).Decode(&event)
	logrus.Info(event)
	event, err = Database.UpdateEvent(event.Id, event.Title, event.Date)
	if err != nil {
		rw.WriteHeader(500)
		logrus.Error("DB error :" + err.Error())
		return
	}

	event.Date = strings.Replace(event.Date, "T", " ", 10)
	response, err := json.Marshal(event)
	if err != nil {
		rw.WriteHeader(500)
		logrus.Error("Json marshall error :" + err.Error())
		return
	}

	rw.WriteHeader(200)
	_, err = rw.Write(response)
	if err != nil {
		logrus.Error("error :" + err.Error())
		return
	}
}

func DeleteEvent(rw http.ResponseWriter, r *http.Request) {

}

func EventForDay(rw http.ResponseWriter, r *http.Request) {
	result, err := Database.EventForDay()
	if err != nil {
		rw.WriteHeader(500)
		logrus.Error("DB error :" + err.Error())
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		rw.WriteHeader(500)
		logrus.Error("Json marshall error :" + err.Error())
		return
	}

	rw.WriteHeader(200)
	_, err = rw.Write(response)
	if err != nil {
		logrus.Error("error :" + err.Error())
		return
	}
}
