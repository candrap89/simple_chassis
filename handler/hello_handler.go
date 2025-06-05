package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chassis/go-chassis/server/restful"
)

// HelloService struct implements the URLPatterns
type HelloService struct{}

// URLPatterns maps routes to handlers
func (s *HelloService) URLPatterns() []restful.Route {
	return []restful.Route{
		{
			Method: http.MethodGet,
			Path:   "/hello",
			Func:   HelloHandler,
		},
	}
}

// HelloHandler writes a simple text response
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Chassis from a separate handler!")
}
