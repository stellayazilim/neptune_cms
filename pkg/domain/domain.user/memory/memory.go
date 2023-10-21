package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	domain_account "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	"github.com/stellayazilim/neptune_cms/pkg/storage/memory"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

type memoryRepository struct {
	users *map[uuid.UUID]*aggregates.User
	sync.Mutex
}

func New() domain_account.IUserRepository {
	return &memoryRepository{
		users: memory.Users,
	}
}

func (m *memoryRepository) Create(user aggregates.User) error {

	if ok := user.GetAccount().ID.Valid; !ok {

		user.GetAccount().ID.UUID = uuid.New()
	}
	// extra check if uuid conflict
	if _, ok := (*m.users)[user.GetAccount().ID.UUID]; ok {

		return domain_account.UserAreadyExistsError
	}

	// check whether user is exist or not
	// if exist return error
	if err := func() error {
		for _, u := range *m.users {
			if u.GetAccount().Email == user.GetAccount().Email {
				return domain_account.UserAreadyExistsError
			}
		}
		return nil
	}(); err != nil {
		return err
	}

	m.Lock()
	(*m.users)[user.GetAccount().ID.UUID] = &user
	m.Unlock()

	return nil
}

func (m *memoryRepository) GetAll() ([]aggregates.User, error) {

	users := make([]aggregates.User, 0)
	m.Lock()
	for _, account := range *m.users {
		users = append(users, *account)
	}
	m.Unlock()
	return users, nil
}

func (m *memoryRepository) GetById(id uuid.UUID) (aggregates.User, error) {

	if _, ok := (*m.users)[id]; !ok {
		return aggregates.NewUser(), domain_account.UserNotFoundError
	}

	m.Lock()

	account := (*m.users)[id]
	m.Unlock()
	return *account, nil
}

func (m *memoryRepository) GetByEmail(email value_objects.Email) (aggregates.User, error) {

	m.Lock()
	for _, account := range *m.users {

		if account.GetAccount().Email == email {

			m.Unlock()
			return *account, nil
		}

	}
	m.Unlock()
	return *new(aggregates.User), domain_account.UserNotFoundError
}

func (m *memoryRepository) UpdateById(id uuid.UUID, account aggregates.User) error {

	if _, ok := (*m.users)[id]; !ok {
		return domain_account.UserNotFoundError
	}

	m.Lock()
	(*m.users)[id] = &account
	m.Unlock()
	return nil
}

func (m *memoryRepository) DeleteById(id uuid.UUID) error {

	if _, ok := (*m.users)[id]; !ok {
		return domain_account.UserNotFoundError
	}
	m.Lock()
	delete(*m.users, id)
	m.Unlock()
	return nil
}
