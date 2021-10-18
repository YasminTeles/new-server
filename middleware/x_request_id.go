package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

type XRequestID struct {
	value  string
	header string
}

func NewXRequestID() *XRequestID {
	return &XRequestID{
		value:  uuid.New().String(),
		header: "X-Request-ID",
	}
}

func (requestID *XRequestID) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	requestID.setHeader(response, request)

	next(response, request)
}

func (requestID *XRequestID) setHeader(response http.ResponseWriter, request *http.Request) {
	if request.Header.Get(requestID.header) == "" {
		request.Header.Set(requestID.header, requestID.value)
	}

	response.Header().Set(requestID.header, request.Header.Get(requestID.header))
}
