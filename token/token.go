package token

import (
	"aidanwoods.dev/go-paseto"
	"github.com/JPratama7/util/sync"
	"github.com/JPratama7/util/token/option"
	"log"
	"time"
)

type TokenArgs func(token *paseto.Token) error

type Paseto struct {
	tokenPooler *sync.Pool[*paseto.Token]
	publicKey   paseto.V4AsymmetricPublicKey
	privateKey  paseto.V4AsymmetricSecretKey
	parser      paseto.Parser
	option      option.Option
}

func NewPaseto(publicKey paseto.V4AsymmetricPublicKey, privateKey paseto.V4AsymmetricSecretKey, options ...option.OptionArgs) *Paseto {

	if len(options) == 0 {
		options = []option.OptionArgs{
			option.WithIssuer("default_issuer"),
			option.WithSubject("default_subject"),
			option.WithAudience("default_audience"),
			option.WithExpiration(time.Minute),
		}
	}

	var opts option.Option
	for _, opt := range options {
		opt(&opts)
	}

	return &Paseto{
		tokenPooler: sync.NewPool(func() *paseto.Token {
			token := paseto.NewToken()
			return &token
		}),

		publicKey:  publicKey,
		privateKey: privateKey,
		option:     opts,
		parser:     paseto.NewParser(),
	}
}

func (p *Paseto) Encrypt(options ...TokenArgs) (string, error) {

	token := p.tokenPooler.Get()
	defer p.tokenPooler.Put(token)

	switch len(options) {
	case 0:
		now := time.Now()
		token.SetIssuer(p.option.Issuer)
		token.SetAudience(p.option.Audience)
		token.SetSubject(p.option.Subject)
		token.SetExpiration(now.Add(p.option.Expiration))
		token.SetNotBefore(now)
	default:
		for _, opt := range options {
			log.Println("Applying option")
			err := opt(token)
			if err != nil {
				return "", err
			}
		}
	}

	return token.V4Sign(p.privateKey, nil), nil
}

func (p *Paseto) Decrypt(token string) (*paseto.Token, error) {
	return p.parser.ParseV4Public(p.publicKey, token, nil)
}

type Token interface {
	Decrypt(token string) (*paseto.Token, error)
	Encrypt(options ...TokenArgs) (string, error)
}
