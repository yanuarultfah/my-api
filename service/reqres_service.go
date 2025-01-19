package service

import (
	"my-api/domain/thirdparty"
	"my-api/utils/errors"
)

var (
	ReqresService reqresServiceInterface = &reqresService{}
)

type reqresService struct{}

type reqresServiceInterface interface {
	GetListUser() (*thirdparty.UserListReqRest, *errors.RestErr)
}

func (s *reqresService) GetListUser() (*thirdparty.UserListReqRest, *errors.RestErr) {
	dao := &thirdparty.UserListReqRest{}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}
