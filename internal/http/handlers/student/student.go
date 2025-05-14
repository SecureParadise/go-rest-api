package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/SecureParadise/students-api/internal/types"
	"github.com/SecureParadise/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

// this will return

// router := http.NewServeMux()
// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Welcome to student api"))
// })

// dependency will be injected here
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// validate request, follow zero trust policy
		// use package /go-playground/validator
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		// decode information //searilize
		// w.Write([]byte("Welcome to students api")).
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
