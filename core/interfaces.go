package bots

import (
	"net/http"
	"time"
)

type BotPlatform interface {
	Id() string
	Version() string
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Criticalf(format string, args ...interface{})
}

type BotHost interface {
	GetLogger(r *http.Request) Logger
	GetHttpClient(r *http.Request) *http.Client
	GetBotChatStore(platform string, r *http.Request) BotChatStore
}

type BotContext struct { // TODO: Rename to BotWebhookContext or just WebhookContext (replace old one)
	BotHost BotHost
	BotSettings BotSettings
}

type WebhookHandler interface {
	RegisterHandlers(pathPrefix string, notFound func(w http.ResponseWriter, r *http.Request))
	HandleWebhookRequest(w http.ResponseWriter, r *http.Request)
	GetBotContextAndInputs(r *http.Request) (botContext BotContext, entriesWithInputs []EntryInputs, err error)
	CreateWebhookContext(r *http.Request, botContext BotContext, webhookInput WebhookInput) WebhookContext //TODO: Can we get rid of http.Request? Needed for botHost.GetHttpClient()

	//ProcessInput(input WebhookInput, entry *WebhookEntry)
}

type WebhookEntry interface {
	GetID() int64
	GetTime() time.Time
}

type WebhookInputType int

const (
	WebhookInputUnknown WebhookInputType = iota
	WebhookInputMessage
	WebhookInputPostback
	WebhookInputDelivery
	WebhookInputAttachment
	WebhookInputInlineQuery // Telegram only?
	WebhookInputChoosenInlineResult // Telegram only?
)

type WebhookInput interface { // '/entry/messaging' for Facebook
	GetSender() WebhookSender
	GetRecipient() WebhookRecipient
	GetTime() time.Time

	InputType() WebhookInputType
	InputMessage() WebhookMessage
	InputPostback() WebhookPostback
	InputDelivery() WebhookDelivery
}

type WebhookActor interface {
	GetID() int64
}

type WebhookSender interface {
	WebhookActor
}

type WebhookRecipient interface {
	WebhookActor
}

type WebhookMessage interface {
	IntID() int64
	StringID() string
	Sequence() int // 'seq' for Facebook, '???' for Telegram
	Text() string
}

type WebhookPostback interface {
	Payload() string
}

type WebhookDelivery interface {
	Payload() string
}

type WebhookAttachment interface {
	Type() string       // Enum(image, video, audio) for Facebook
	PayloadUrl() string // 'payload.url' for Facebook
}

type WebhookResponder interface {
	SendMessage(m MessageFromBot) error
}

type InputMessage interface {
	Text() string
}