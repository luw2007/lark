package lark

import (
	"context"
	"net/http"
	"strings"
	"time"
)

// New new Lark client
func New(options ...ClientOptionFunc) *Lark {
	return newClient("", options)
}

// AppLink ref Lark.AppLink instance
var AppLink = New().AppLink

// ClientOptionFunc new Lark client option
type ClientOptionFunc func(*Lark)

// WithAppCredential set app credential
func WithAppCredential(appID, appSecret string) ClientOptionFunc {
	return func(r *Lark) {
		r.appID = appID
		r.appSecret = appSecret
	}
}

// WithCustomBot set custom bot
func WithCustomBot(customBotWebHookURL, customBotSecret string) ClientOptionFunc {
	return func(r *Lark) {
		r.customBotWebHookURL = customBotWebHookURL
		r.customBotSecret = customBotSecret
	}
}

// WithEventCallbackVerify set event callback verify
func WithEventCallbackVerify(encryptKey, verificationToken string) ClientOptionFunc {
	return func(r *Lark) {
		r.encryptKey = encryptKey
		r.verificationToken = verificationToken
	}
}

// WithHelpdeskCredential set helpdesk credential
func WithHelpdeskCredential(helpdeskID, helpdeskToken string) ClientOptionFunc {
	return func(r *Lark) {
		r.helpdeskID = helpdeskID
		r.helpdeskToken = helpdeskToken
	}
}

// WithTimeout set timeout
func WithTimeout(timeout time.Duration) ClientOptionFunc {
	return func(r *Lark) {
		r.timeout = timeout
	}
}

// WithHttpClient set http client
func WithHttpClient(cli HttpClient) ClientOptionFunc {
	return func(r *Lark) {
		r.httpClient = cli
	}
}

// WithLogger set logger
func WithLogger(logger Logger, level LogLevel) ClientOptionFunc {
	return func(lark *Lark) {
		lark.logger = logger
		lark.logLevel = level
	}
}

// WithISV set isv
func WithISV(isISV bool) ClientOptionFunc {
	return func(lark *Lark) {
		lark.isISV = isISV
	}
}

// WithStore set store
func WithStore(store Store) ClientOptionFunc {
	return func(lark *Lark) {
		lark.store = store
	}
}

// WithOpenBaseURL set open base url
func WithOpenBaseURL(baseURL string) ClientOptionFunc {
	return func(lark *Lark) {
		lark.openBaseURL = strings.TrimRight(baseURL, "/")
	}
}

// WithWWWBaseURL set www base url
func WithWWWBaseURL(baseURL string) ClientOptionFunc {
	return func(lark *Lark) {
		lark.wwwBaseURL = strings.TrimRight(baseURL, "/")
	}
}

// MethodOptionFunc new method option
type MethodOptionFunc func(*MethodOption)

// WithUserAccessToken set user access token
func WithUserAccessToken(token string) MethodOptionFunc {
	return func(option *MethodOption) {
		option.userAccessToken = token
	}
}

// MethodOption method option
type MethodOption struct {
	userAccessToken string
}

func newMethodOption(options []MethodOptionFunc) *MethodOption {
	opt := new(MethodOption)
	for _, v := range options {
		v(opt)
	}
	return opt
}

func newClient(tenantKey string, options []ClientOptionFunc) *Lark {
	r := &Lark{
		timeout:      time.Second * 3,
		isISV:        false,
		tenantKey:    tenantKey,
		store:        NewStoreMemory(),
		mock:         new(Mock),
		eventHandler: new(eventHandler),
		openBaseURL:  "https://open.feishu.cn",
		wwwBaseURL:   "https://www.feishu.cn",
	}
	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	if r.httpClient == nil {
		r.httpClient = newDefaultHttpClient(r.timeout) // 这个时候之前 timeout 可能还没有设置
	}

	r.initService()

	return r
}

// HttpClient http client iface
type HttpClient interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}
