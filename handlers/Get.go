package handlers

import (
	"net/http"

	"github.com/Chara-X/state"
)

type Get struct{ Store state.MemoryStore[entry] }

func (g *Get) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var key = request.URL.Query().Get("key")
	if entry, ok := g.Store.Load(key); ok {
		writer.Header().Set("ETag", entry.eTag)
		writer.Write(entry.value)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

type entry struct {
	value []byte
	eTag  string
}
