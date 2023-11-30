package Models

type AggregateRoot[TEntity any] struct {
	Root TEntity
}
