package handles

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/carlosdamazio/k6-test/internal/rest/entities"
	"github.com/carlosdamazio/k6-test/internal/services/message"
)

const (
	baseURI = "messages"
	URISep  = "/"
)

type MessageHandler interface {
	MessageHandle(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	svc message.MessageService
}

func NewMessageHandle(svc message.MessageService) *Handler {
	return &Handler{svc}
}

func (h *Handler) MessageHandle(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case "GET":
		args := strings.Split(r.RequestURI, URISep)[1:]

		switch len(args) {
		case 1:
			messages := h.svc.List()
			encodedMessages, err := json.MarshalIndent(messages, "", "    ")
			if err != nil {
				goto ERROR
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", string(encodedMessages))
		case 2:
			message, err := h.svc.Get(args[1])
			if err != nil {
				goto ERROR
			}
			encodedMessage, err := json.MarshalIndent(message, "", "    ")
			if err != nil {
				goto ERROR
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", string(encodedMessage))
		default:
			err = errors.New("get: invalid path")
			goto ERROR
		}
	case "POST":
		var (
			reqData *entities.PostMessageRequest
			body    []byte
		)

		defer r.Body.Close()

		if body, err = io.ReadAll(r.Body); err != nil {
			goto ERROR
		}

		if len(body) == 0 {
			err = errors.New("get: request body is empty")
			goto ERROR
		}

		if err = json.Unmarshal(body, &reqData); err != nil {
			goto ERROR
		}

		instance := h.svc.Post(reqData)
		encodedInstance, err := json.MarshalIndent(instance, "", "    ")
		if err != nil {
			goto ERROR
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", string(encodedInstance))
	case "PATCH":

	}

	return
ERROR:
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "server error: %v", err)
}
