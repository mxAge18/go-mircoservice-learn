package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var JsonContentType string = "application/json"

type UserTransporter interface {
	EncodeRequestFunc(ctx context.Context, r *http.Request, i interface{}) error
	DecodeResponseFunc(ctx context.Context, res *http.Response) (response interface{}, err error)
}

type userTransport struct {
}

func NewUserTransporter() UserTransporter {
	return &userTransport{}
}

func (u *userTransport) EncodeRequestFunc(ctx context.Context, r *http.Request, i interface{}) error {
	userRequest := i.(UserRequest)
	r.URL.Path += "/user/" + strconv.Itoa(userRequest.UserId)
	return nil
}

func (u *userTransport) DecodeResponseFunc(ctx context.Context, res *http.Response) (response interface{}, err error) {
	if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("error status code")
	}
	var userResponse UserResponse
	err = json.NewDecoder(res.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}
