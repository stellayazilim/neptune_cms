package Models

type Entity[TValueObject any] struct {
	ID TValueObject
}

func (e *Entity[TId]) GetID() TId {
	return e.ID
}
