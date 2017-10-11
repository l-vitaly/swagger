package swagger

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type swaggerHandler struct {
	filepath    string
	allowOrigin string
}

func (h *swaggerHandler) writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
}

func (h *swaggerHandler) getContentTypeByExt(ext string) string {
	switch ext {
	default:
		return "application/yaml"
	case ".json":
		return "application/json"
	}
}

func (h *swaggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(h.filepath)
	if err != nil {
		h.writeError(w, err)
		return
	}
	data, err := ioutil.ReadFile(h.filepath)
	if err != nil {
		h.writeError(w, err)
		return
	}
	w.Header().Set("Content-type", h.getContentTypeByExt(filepath.Ext(path)))
	w.Header().Set("Access-Control-Allow-Origin", h.allowOrigin)
	w.Write(data)
}

// Handler swagger file handler
func Handler(filepath, allowOrigin string) http.Handler {
	return &swaggerHandler{
		filepath:    filepath,
		allowOrigin: allowOrigin,
	}
}
