package value_objects

type Email string

func NewEmail(s string) *Email {
	e := Email(s)
	return &e
}
