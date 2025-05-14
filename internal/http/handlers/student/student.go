package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/SecureParadise/students-api/internal/types"
	"github.com/SecureParadise/students-api/internal/utils/response"
)

// this will return

// router := http.NewServeMux()
// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Welcome to student api"))
// })

// dependency will be injected here
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("creatinga student")
		// decode information //searilize
		// w.Write([]byte("Welcome to students api")).
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
