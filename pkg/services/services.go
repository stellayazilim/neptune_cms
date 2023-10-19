package services

type ServiceConfig[T interface{}] func(*T) error
