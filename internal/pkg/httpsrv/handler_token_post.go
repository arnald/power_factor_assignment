package httpsrv

import (
	"net/http"
)

func (s *Server) handlerTokenPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST request successful"))
}
