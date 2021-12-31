package customhandler

import (
	"fmt"
	"net/http"
)

type msghandler struct {
	message string
}

func (m *msghandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}
