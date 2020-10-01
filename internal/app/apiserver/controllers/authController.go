package controllers

import (
	"encoding/json"
	"net/http"
)

// Account ...
type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWTToken ...
type JWTToken struct {
	Token string `json:"token"`
}

// HandleLogin ...
func HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		account := &Account{}
		err := json.NewDecoder(r.Body).Decode(account)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			message := message(false, "Invalid request")
			json.NewEncoder(w).Encode(message)
			return
		}

		if account.Email == "" {
			w.WriteHeader(http.StatusForbidden)
			message := message(false, "Email is required")
			json.NewEncoder(w).Encode(message)
			return
		}

		if account.Password == "" {
			w.WriteHeader(http.StatusForbidden)
			message := message(false, "Password is required")
			json.NewEncoder(w).Encode(message)
			return
		}

		token := JWTToken{
			Token: "123qweasdzxc456rtyfghvbn",
		}

		json.NewEncoder(w).Encode(&token)
	}
}

func message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
