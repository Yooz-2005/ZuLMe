package server

import (
	"context"
	"user_srv/internal/logic"
	user "user_srv/proto_user"
)

type ServerUser struct {
	user.UnimplementedUserServer
}

func (s ServerUser) UserRegister(ctx context.Context, in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	register, err := logic.UserRegister(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

func (s ServerUser) SendCode(ctx context.Context, in *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	code, err := logic.SendCode(in)
	if err != nil {
		return nil, err
	}
	return code, nil
}
