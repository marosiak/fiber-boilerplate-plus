package sessioncontext

import (
	"github.com/gofiber/fiber/v2"
	fiberSession "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

const SessionPropsKey = "sessionProps"

type sessionProps map[string]any
type SessionContext struct {
	log   *logrus.Entry
	store *fiberSession.Store
}

func New(store *fiberSession.Store, log *logrus.Entry) *SessionContext {
	store.RegisterType(sessionProps{})
	return &SessionContext{store: store, log: log}
}

func (p *SessionContext) ClearAll(c *fiber.Ctx) {
	session, err := p.store.Get(c)
	if err != nil {
		p.log.WithError(err).Error("clearing SessionProps failed")
		return
	}

	session.Set(SessionPropsKey, nil)
	err = session.Save()
	if err != nil {
		p.log.WithError(err).Error("saving SessionProps failed")
		return
	}
}

func (p *SessionContext) Get(c *fiber.Ctx, key string, defaultValue any) any {
	log := p.log.WithField("key", key)
	session, err := p.store.Get(c)
	if err != nil {
		log.WithError(err).Error("getting session in SessionContext")
		return nil
	}

	props := session.Get(SessionPropsKey)
	if props == nil {
		return defaultValue
	}

	val, ok := props.(sessionProps)[key]
	if !ok {
		return defaultValue
	}
	return val
}

func (p *SessionContext) Set(c *fiber.Ctx, key string, value any) {
	log := p.log.WithFields(logrus.Fields{
		"key": key,
		"val": value,
	})
	session, err := p.store.Get(c)
	if err != nil {
		log.WithError(err).Error("pushing SessionProp failed")
		return
	}

	sesProps := session.Get(SessionPropsKey)
	if sesProps == nil {
		sesProps = make(sessionProps)
	}

	props := sesProps.(sessionProps)
	props[key] = value

	session.Set(SessionPropsKey, props)
	err = session.Save()
	if err != nil {
		log.WithError(err).Error("saving SessionProp failed in SessionContext")
		return
	}
}
