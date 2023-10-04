package router

import (
	"context"
	"fmt"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/commands/demo"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/loader"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"runtime/debug"
	"sync"
)

type Router struct {
	bot     *tgbotapi.BotAPI
	routes  map[string]commander
	updates tgbotapi.UpdatesChannel
	cancel  context.CancelFunc
	proc    *loader.Single
	once    sync.Once
}

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
	r := &Router{
		bot:    bot,
		routes: map[string]commander{},
	}
	r.Route("demo", "subdomain", demo.NewDemoCommander(bot))
	return r
}

func (r *Router) getRoute(domain, subdomain string) commander {
	return r.routes[domain+"/"+subdomain]
}

func (r *Router) Route(domain, subdomain string, commander commander) {
	if c, ok := commander.(configurableCommander); ok {
		c.Config(commanderCfg{
			BotAPI:    r.bot,
			Domain:    domain,
			Subdomain: subdomain,
		})
	}
	r.routes[domain+"/"+subdomain] = commander
}

// HandleUpdate You should not use this method directly. This is left for
// backward compatibility only. Use the Start/Close methods to start and
// stop the router.
func (r *Router) HandleUpdate(update tgbotapi.Update) {
	const op = "Router.HandleUpdate"

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Error().Msgf("%s: recovered from panic: %v\n%v", op, panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		r.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		r.handleMessage(update.Message)
	}
}

func (r *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	const op = "Router.handleCallback"

	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("%s: error parsing callback data `%s` - %v", op, callback.Data, err)
		return
	}

	route := r.getRoute(callbackPath.Domain, callbackPath.Subdomain)
	if route == nil {
		log.Printf("%s: unknown callback path - %s/%s", op, callbackPath.Domain, callbackPath.Subdomain)
		return
	}

	route.HandleCallback(callback, callbackPath)
}

func (r *Router) handleMessage(msg *tgbotapi.Message) {
	const op = "Router.handleMessage"

	if !msg.IsCommand() {
		r.showCommandFormat(msg)
		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("%s: error parsing command `%s` - %v", op, msg.Command(), err)
		r.showCommandFormat(msg)
		return
	}

	route := r.getRoute(commandPath.Domain, commandPath.Subdomain)
	if route == nil {
		log.Printf("%s: unknown command path - %s/%s", op, commandPath.Domain, commandPath.Subdomain)
		r.send(msg.Chat.ID, "Unknown command path: %s", msg.Text)
		return
	}

	route.HandleCommand(msg, commandPath)
}

func (r *Router) showCommandFormat(msg *tgbotapi.Message) {
	r.send(msg.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")
}

func (r *Router) send(chatID int64, msg string, a ...any) {
	const op = "Router.send"

	output := tgbotapi.NewMessage(chatID, fmt.Sprintf(msg, a...))

	if _, err := r.bot.Send(output); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}

func (r *Router) Run(ctx context.Context) {
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			return
		case update := <-r.updates:
			r.HandleUpdate(update)
		}
	}
}

func (r *Router) Start(ctx context.Context, cfg tgbotapi.UpdateConfig) (err error) {
	r.once.Do(func() {
		err = r.start(ctx, cfg)
	})
	return err
}

func (r *Router) start(ctx context.Context, cfg tgbotapi.UpdateConfig) error {
	const op = "Router.Start"

	updates, err := r.bot.GetUpdatesChan(cfg)
	if err != nil {
		return fmt.Errorf("%s: can't get bot update channel: %w", op, err)
	}

	ctx, cancel := context.WithCancel(ctx)
	r.updates = updates
	r.cancel = cancel

	r.proc = loader.Start(ctx, r)

	return nil
}

func (r *Router) Close() {
	r.cancel()
	r.proc.Wait()
}
