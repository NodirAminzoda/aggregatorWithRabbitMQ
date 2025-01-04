package service

import (
	"aggreagtor/internal/domain/precheck"
	"aggreagtor/internal/modules/adapter"
	"aggreagtor/utils"
	"context"
)

type Service struct {
	adapter adapter.IAdapter
}

func New(adapter adapter.IAdapter) *Service {
	return &Service{
		adapter: adapter,
	}
}

type IService interface {
	PreCheckService(ctx context.Context, req precheck.Request) (*precheck.Response, error)
}

func (s *Service) PreCheckService(ctx context.Context, req precheck.Request) (*precheck.Response, error) {
	// Проверка аккаунта перед выполнением запроса
	if !precheck.ValidateAccount(req.Account) {
		return nil, precheck.InvalidAccount
	}

	// Вызов метода PreCheck адаптера
	response, err := s.adapter.PreCheck(req, utils.GetRequestID(ctx))
	if err != nil {
		return nil, err
	}

	return response, nil
}
