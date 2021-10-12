package middleware

import (
	"net/http"

  "github.com/google/uuid"
)

type XRequestID struct {
  value string
  header string
}

func NewXRequestID() *XRequestID{
  return &XRequestID{
    value: uuid.New().String(),
    header: "X-Request-ID",
  }
}

func (requestID *XRequestID) ServeHTTP(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  if req.Header.Get(requestID.header) == "" {
    req.Header.Set(requestID.header, requestID.value)
  }

  res.Header().Set(requestID.header,req.Header.Get(requestID.header))

  next(res, req)
}
