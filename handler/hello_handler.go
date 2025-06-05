package handler

import (
	"context"
	"net/http"
)

// HelloService implements the REST schema
type HelloService struct{}

// Hello handles GET requests on /hello
func (h *HelloService) Hello(ctx context.Context, rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello from Go-Chassis with URL pattern!"))
}
