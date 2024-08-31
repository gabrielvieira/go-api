package service

import "context"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	return nil
}
