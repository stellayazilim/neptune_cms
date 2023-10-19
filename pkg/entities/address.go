package entities

import "github.com/stellayazilim/neptune_cms/pkg/value_objects"

type Address struct {
	Base
	ID           uint64
	Country      string
	State        string
	City         string
	Province     string
	AddressLines value_objects.AddressLines
	No           string
	PostCode     string
}
