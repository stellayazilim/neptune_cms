package MeCommand

import "github.com/stellayazilim/neptune.domain/User"

type (
	MeCommand struct {
		Email string
	}

	MeCommandResponse struct {
		User User.UserAggregate
	}
)
