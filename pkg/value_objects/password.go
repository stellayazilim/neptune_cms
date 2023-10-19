package value_objects

type Password []byte

func (p *Password) ToString() string {
	return string(*p)
}

func NewPassword(s string) Password {
	p := Password(s)

	return p
}
