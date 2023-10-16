package value_objects

type Password []byte

func (p *Password) ToString() string {
	return string(*p)
}

func NewPassword(b []byte) *Password {
	p := new(Password)
	*p = b
	return p
}
