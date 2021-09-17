package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var JsonContentType string = "application/json"

type Transporter interface {
	EncodeRequestFunc(ctx context.Context, r *http.Request, i interface{}) error
	DecodeResponseFunc(ctx context.Context, res *http.Response) (response interface{}, err error)
}

type transport struct {
}

func NewUserTransporter() Transporter {
	return &transport{}
}

func (u *transport) EncodeRequestFunc(ctx context.Context, r *http.Request, i interface{}) error {
	userRequest := i.(Request)
	r.URL.Path += "/user/" + strconv.Itoa(userRequest.UserId)
	return nil
}

func (u *transport) DecodeResponseFunc(ctx context.Context, res *http.Response) (response interface{}, err error) {
	if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("error status code")
	}
	var userResponse Response
	err = json.NewDecoder(res.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}
