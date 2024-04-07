package server

import (
	"io"
	"net/http"
	"time"

	"github.com/Chara-X/state"
	"github.com/google/uuid"
)

type Handler struct{ store state.MemoryStore[entry] }

func New() http.Handler { return &Handler{state.NewMemoryStore[entry]()} }
func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var key = request.URL.Query().Get("key")
	var multiplexer = http.NewServeMux()
	multiplexer.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		if entry, ok := handler.store.Get(key); ok {
			writer.Header().Set("ETag", entry.etag)
			writer.Write(entry.value)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	})
	multiplexer.HandleFunc("/set", func(writer http.ResponseWriter, request *http.Request) {
		var value, _ = io.ReadAll(request.Body)
		var duration, _ = time.ParseDuration(request.URL.Query().Get("duration"))
		if etag := request.Header.Get("If-Match"); etag != "" {
			if entry, ok := handler.store.Get(key); ok && etag != entry.etag {
				writer.WriteHeader(http.StatusPreconditionFailed)
				return
			}
		}
		handler.store.Set(key, entry{value, uuid.NewString()}, duration)
	})
	multiplexer.ServeHTTP(writer, request)
}

type entry struct {
	value []byte
	etag  string
}
