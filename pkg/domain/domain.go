package domain

type RepositoryConfig[c any] func(repository *c) error
