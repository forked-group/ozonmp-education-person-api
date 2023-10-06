package interfaces

import (
	"context"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/path"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type ConfigurableCommander interface {
	Commander
	Config(cfg CommanderCfg)
}

type CommanderCfg struct {
	BotAPI    *tgbotapi.BotAPI
	Domain    string
	Subdomain string
}

type PersonCommander interface {
	ConfigurableCommander
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)  // return error not implemented - where? how?
	Edit(inputMsg *tgbotapi.Message) // return error not implemented - where? how?
}

type PersonService interface {
	Describe(personID uint64) (*education.Person, error)
	List(cursor uint64, limit uint64) ([]education.Person, error)
	Create(create education.PersonCreate) (uint64, error)
	Update(personID uint64, person education.PersonCreate) (bool, error)
	Remove(personID uint64) (bool, error)
}

type PersonRepo interface {
	DescribePerson(ctx context.Context, personID uint64) (*education.Person, error)
	ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]education.Person, error)
	CreatePerson(ctx context.Context, person education.PersonCreate) (uint64, error)
	UpdatePerson(ctx context.Context, personID uint64, person education.PersonCreate) (bool, error)
	RemovePerson(ctx context.Context, personID uint64) (bool, error)
}

type PersonEventRepo interface {
	Lock(ctx context.Context, n uint64) ([]education.PersonEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) (uint64, error)
	Remove(ctx context.Context, eventIDs []uint64) (uint64, error)
}
