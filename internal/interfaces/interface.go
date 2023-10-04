package interfaces

import (
	"context"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/router"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type PersonService interface {
	Describe(personID uint64) (*education.Person, error)
	List(cursor uint64, limit uint64) ([]education.Person, error)
	Create(education.Person) (uint64, error)
	Update(personID uint64, person education.Person) (bool, error)
	Remove(personID uint64) (bool, error)
}

// PersonRepo is DAO for Person
type PersonRepo interface {
	DescribePerson(ctx context.Context, personID uint64) (*education.Person, error)
	ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]education.Person, error)
	CreatePerson(ctx context.Context, person education.Person) (uint64, error)
	UpdatePerson(ctx context.Context, personID uint64, person education.Person) (bool, error)
	RemovePerson(ctx context.Context, personID uint64) (bool, error)
}

type PersonCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented - where? how?
	Edit(inputMsg *tgbotapi.Message) // return error not implemented - where? how?

	router.Commander
}
