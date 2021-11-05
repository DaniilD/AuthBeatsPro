package operations

import "net/http"

type Operation interface {
	Handle(r *http.Request) (interface{}, error)
}
