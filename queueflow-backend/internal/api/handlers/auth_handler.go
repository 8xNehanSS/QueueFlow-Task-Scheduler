package handlers

import (
	"encoding/json"
	"net/http"

	"queueflow/internal/auth"
)

type AuthHandler struct {
	service *auth.AuthService
}

func NewAuthHandler(
	service *auth.AuthService,
) *AuthHandler {

	return &AuthHandler{
		service: service,
	}

}

func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {

	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).
		Decode(&body)

	err :=
		h.service.Register(
			body.Username,
			body.Email,
			body.Password,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			400,
		)

		return
	}

	json.NewEncoder(w).
		Encode(
			map[string]string{
				"message": "registered",
			},
		)

}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {

	var body struct {
		Email string `json:"email"`

		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).
		Decode(&body)

	token, err :=
		h.service.Login(
			body.Email,
			body.Password,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			401,
		)

		return
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   86400,
		},
	)

	json.NewEncoder(w).
		Encode(
			map[string]string{
				"login": "success",
			},
		)

}
