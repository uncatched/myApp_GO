package controllers

import (
	"encoding/json"
	"my-api/internal/app/apiserver/models"
	"my-api/internal/validate"
	"net/http"
)

// HandleLogin ...
func HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		account := &models.AuthUser{}
		err := json.NewDecoder(r.Body).Decode(account)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			message := message(false, err.Error())
			json.NewEncoder(w).Encode(message)
			return
		}

		v := validate.Validator{}
		err = v.Validate(account)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			message := message(false, err.Error())
			json.NewEncoder(w).Encode(message)
			return
		}

		token := models.JWTToken{
			Token: "123qweasdzxc456rtyfghvbn",
		}

		json.NewEncoder(w).Encode(&token)
	}
}

// HandleSignup ...
func HandleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		account := &models.AuthUser{}
		err := json.NewDecoder(r.Body).Decode(account)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			message := message(false, "Invalid request")
			json.NewEncoder(w).Encode(message)
			return
		}

		v := validate.Validator{}
		err = v.Validate(account)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			message := message(false, err.Error())
			json.NewEncoder(w).Encode(message)
			return
		}

		token := models.JWTToken{
			Token: "123qweasdzxc456rtyfghvbn",
		}

		json.NewEncoder(w).Encode(&token)
	}
}

func message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
