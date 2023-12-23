package Models

import "reflect"

type ValueObject[T any] struct {
	value T
}

func (v *ValueObject[T]) GetValue() T {
	return v.value
}

func (v *ValueObject[T]) IsEqual(right ValueObject[T]) bool {
	return reflect.DeepEqual(v, &right)
}

func NewValueObject[T any](value T) ValueObject[T] {

	return ValueObject[T]{
		value: value,
	}
}
