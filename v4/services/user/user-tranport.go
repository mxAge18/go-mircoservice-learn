package user

import (
	"context"
	"encoding/json"
	"errors"
	httpTransport "github.com/go-kit/kit/transport/http"
	routerMux "github.com/gorilla/mux"
	"go-mircoservice-learn/v4/utils"
	"net/http"
	"strconv"
)

var JsonContentType string = "application/json"

type UserTransporter interface {
	DecodeRequest(c context.Context, r *http.Request) (decodeRes interface{}, err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, request interface{}) (err error)
	GetOptions() []httpTransport.ServerOption
}

type userTransport struct {
	options []httpTransport.ServerOption
}

func NewUserTransporter() UserTransporter {
	t := &userTransport{}
	t.options = []httpTransport.ServerOption{
		httpTransport.ServerErrorEncoder(t.CustomErrorEncoder),
	}
	return t
}

func (u *userTransport) GetOptions() []httpTransport.ServerOption {
	return u.options
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

func (u *userTransport) CustomErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-Type", contentType)

	code := http.StatusInternalServerError
	if cusErr, ok := err.(*utils.CustomErr); ok {
		code = cusErr.Code
	}
	w.WriteHeader(code)
	w.Write(body)
}
