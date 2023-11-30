package User

import (
	. "github.com/stellayazilim/neptune_cms/internal/Domain/Common/Models"
	. "github.com/stellayazilim/neptune_cms/internal/Domain/User/Entities"
)

type UserAggregate struct {
	AggregateRoot[UserEntity]
}

func CreateUserAggregate() UserAggregate {
	return UserAggregate{}
}

func EmptyUserAggregate() UserAggregate {
	return UserAggregate{
		AggregateRoot: AggregateRoot[UserEntity]{
			Root: *new(UserEntity),
		},
	}
}
