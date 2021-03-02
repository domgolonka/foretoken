package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/domgolonka/foretoken/app/services"
	"github.com/domgolonka/foretoken/lib/parse"
	"github.com/sirupsen/logrus"
)

type ServiceData struct {
	Result interface{} `json:"result"`
}

type ServiceErrors struct {
	Errors services.FieldErrors `json:"errors"`
}

type RequestError struct {
	Error string `json:"error"`
}

func WriteData(w http.ResponseWriter, httpCode int, str *[]string) {
	stringByte := strings.Join(*str, "\x0A") // x20 = space and x00 = null

	w.WriteHeader(httpCode)
	_, err := w.Write([]byte(stringByte))
	if err != nil {
		WriteErrors(w, err)
	}
}

func WriteErrors(w http.ResponseWriter, err error) {
	switch x := err.(type) {
	case services.FieldErrors:
		WriteJSON(w, http.StatusUnprocessableEntity, ServiceErrors{Errors: x})
	case parse.Error:
		writeParseErrors(w, x)
	default:
		WriteJSON(w, http.StatusInternalServerError, RequestError{Error: err.Error()})
	}
}

func writeParseErrors(w http.ResponseWriter, err parse.Error) {
	switch err.Code {
	case parse.UnsupportedMediaType:
		WriteJSON(w, http.StatusUnsupportedMediaType, RequestError{Error: err.Message})
	default:
		WriteJSON(w, http.StatusBadRequest, RequestError{Error: err.Message})
	}
}

func WriteNotFound(w http.ResponseWriter, resource string) {
	WriteJSON(w, http.StatusNotFound, ServiceErrors{Errors: services.FieldErrors{{resource, services.ErrNotFound}}}) //nolint
}

func WriteJSON(w http.ResponseWriter, httpCode int, d interface{}) {
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, err = w.Write(j)
	if err != nil {
		logrus.Error(err)
	}
}
