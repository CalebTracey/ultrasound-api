package routes

import (
	"encoding/json"
	"github.com/CalebTracey/ultrasound-api/internal/facade"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Service facade.UltrasoundFacadeI
}

func (h Handler) InitializeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/api/auth/user/{username}", h.Authentication()).Methods(http.MethodGet)

	return r
}

func (h Handler) Authentication() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]

		result, status := h.Service.Authentication(username)
		if len(result.Message.ErrorLog) == 0 {
			result.Message.Status = strconv.Itoa(http.StatusOK)
		}

		result = SetAuthenticationResponse(result)
		_ = json.NewEncoder(writeHeader(rw, status)).Encode(result)

	}
}

func writeHeader(rw http.ResponseWriter, code int) http.ResponseWriter {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	return rw
}
