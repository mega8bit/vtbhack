package main

import (
	"context"
	"github.com/graphql-go/handler"
	"net/http"
	"strings"
)

type AuthByToken struct {
	*handler.Handler
}

func (a *AuthByToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer")
	token = strings.TrimSpace(token)
	ctx := context.WithValue(r.Context(), "token", token)
	a.Handler.ServeHTTP(w, r.WithContext(ctx))
}
