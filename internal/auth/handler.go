package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"myFirstHttpServer/configs"
	"myFirstHttpServer/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if payload.Email == "" || payload.Password == "" {
			res.Json(w, "Email and Pass are required", 402)
			return
		}
		match, err := regexp.MatchString(`[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`, payload.Email)
		if !match {
			res.Json(w, "Wrong email", 402)
			return
		}
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}
		fmt.Println(payload)
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Страница логина")
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Страница регистрации")
	}
}
