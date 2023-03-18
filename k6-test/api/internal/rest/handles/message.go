package handles

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/carlosdamazio/k6-test/internal/services/message"
)

const (
	baseURI = "messages"
	URISep  = "/"
)

type MessageHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	svc message.MessageService
}

func NewMessageHandle(svc message.MessageService) *Handler {
	return &Handler{svc}
}

func (h *Handler) MessageHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		for _, arg := range strings.Split(r.RequestURI, URISep) {
			fmt.Fprintf(w, "arg: %s\n", arg)
		}
	}
}
