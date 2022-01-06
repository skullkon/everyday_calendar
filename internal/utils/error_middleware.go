package utils

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logrus.Info(request.RequestURI)
		next(writer, request)
	})
}
