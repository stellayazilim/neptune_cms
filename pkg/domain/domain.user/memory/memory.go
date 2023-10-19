package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/stellayazilim/neptune_cms/pkg/aggregates"
	domain_account "github.com/stellayazilim/neptune_cms/pkg/domain/domain.user"
	"github.com/stellayazilim/neptune_cms/pkg/value_objects"
)

type memoryRepository struct {
	accounts map[uuid.UUID]aggregates.User
	sync.Mutex
}

func New(accounts map[uuid.UUID]aggregates.User) domain_account.IAccountRepository {

	if accounts != nil {
		return &memoryRepository{
			accounts: accounts,
		}
	}
	return &memoryRepository{
		accounts: make(map[uuid.UUID]aggregates.User, 0),
	}
}

func (m *memoryRepository) Create(user aggregates.User) error {
	fmt.Println(user.GetAccount().ID.Valid)
	if ok := user.GetAccount().ID.Valid; !ok {
		return domain_account.UserInvalidIDError
	}

	if _, ok := m.accounts[user.GetAccount().ID.UUID]; ok {
		return domain_account.UserAreadyExistsError
	}

	m.Lock()
	m.accounts[user.GetAccount().ID.UUID] = user
	m.Unlock()

	return nil
}

func (m *memoryRepository) GetAll() ([]aggregates.User, error) {

	accounts := make([]aggregates.User, 0)
	m.Lock()
	for _, account := range m.accounts {
		accounts = append(accounts, account)
	}
	m.Unlock()
	return accounts, nil
}

func (m *memoryRepository) GetById(id uuid.UUID) (aggregates.User, error) {

	if _, ok := m.accounts[id]; !ok {
		return aggregates.NewUser(), domain_account.UserNotFoundError
	}

	m.Lock()

	account := m.accounts[id]
	m.Unlock()
	return account, nil
}

func (m *memoryRepository) GetByEmail(email value_objects.Email) (aggregates.User, error) {

	m.Lock()

	for _, account := range m.accounts {

		if account.GetAccount().Email == email {
			return account, nil
		}

	}

	m.Unlock()
	return *new(aggregates.User), domain_account.UserNotFoundError
}

func (m *memoryRepository) UpdateById(id uuid.UUID, account aggregates.User) error {

	if _, ok := m.accounts[id]; !ok {
		return domain_account.UserNotFoundError
	}

	m.Lock()
	m.accounts[id] = account
	m.Unlock()
	return nil
}

func (m *memoryRepository) DeleteById(id uuid.UUID) error {

	if _, ok := m.accounts[id]; !ok {
		return domain_account.UserNotFoundError
	}
	m.Lock()
	delete(m.accounts, id)
	m.Unlock()
	return nil
}
