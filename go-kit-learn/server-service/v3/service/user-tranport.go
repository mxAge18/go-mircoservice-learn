package service

import (
	"context"
	"encoding/json"
	"errors"
	routerMux "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var JsonContentType string = "application/json"

type UserTransporter interface {
	DecodeRequest(c context.Context, r *http.Request) (decodeRes interface{}, err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, request interface{}) (err error)
}

type userTransport struct {
}

func NewUserTransporter() UserTransporter {
	return &userTransport{}
}

func (u *userTransport) DecodeRequest(c context.Context, r *http.Request) (decodeRes interface{}, err error) {
	vars := routerMux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		err = errors.New("未解析到参数")
		return
	}

	uid, _ := strconv.Atoi(userId)

	decodeRes = UserRequest{UserId: uid, Method: r.Method}
	return
}

func (u *userTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, request interface{}) (err error) {
	w.Header().Set("Content-Type", JsonContentType)
	return json.NewEncoder(w).Encode(request)
}
