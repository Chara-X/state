package handlers

import (
	"io"
	"net/http"
	"time"

	"github.com/Chara-X/state"
	"github.com/google/uuid"
)

type Set struct{ Store state.MemoryStore[entry] }

func (s *Set) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var key = request.URL.Query().Get("key")
	var value, _ = io.ReadAll(request.Body)
	var duration, _ = time.ParseDuration(request.URL.Query().Get("duration"))
	if eTag := request.Header.Get("If-Match"); eTag != "" {
		if entry, ok := s.Store.Load(key); ok && eTag != entry.eTag {
			writer.WriteHeader(http.StatusPreconditionFailed)
			return
		}
	}
	s.Store.Store(key, entry{value, uuid.NewString()}, duration)
}
