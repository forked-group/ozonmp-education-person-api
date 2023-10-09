package interfaces

import (
	"context"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/path"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
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
	Describe(personID uint64) (*model.Person, error)
	List(cursor uint64, limit uint64) ([]model.Person, error)
	Create(create model.Person) (uint64, error)
	Update(personID uint64, person model.Person, fields model.PersonField) (bool, error)
	Remove(personID uint64) (bool, error)
}

type PersonRepo interface {
	Describe(ctx context.Context, personID uint64) (*model.Person, error)
	List(ctx context.Context, cursor uint64, limit uint64) ([]model.Person, error)
	Create(ctx context.Context, person model.Person) (uint64, error)
	Update(ctx context.Context, personID uint64, person model.Person, fields model.PersonField) (bool, error)
	Remove(ctx context.Context, personID uint64) (bool, error)
}

type PersonEventRepo interface {
	Lock(ctx context.Context, n uint64) ([]model.PersonEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) (uint64, error)
	Remove(ctx context.Context, eventIDs []uint64) (uint64, error)
}
