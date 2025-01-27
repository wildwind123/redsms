// Code generated by ogen, DO NOT EDIT.

package ogencl

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// Login provides Login security value.
	// Имя пользователя redsms.
	Login(ctx context.Context, operationName string) (Login, error)
	// Secret provides Secret security value.
	// Авторизация  \
	// https://docs.redsms.ru/http/getting-started/.
	Secret(ctx context.Context, operationName string) (Secret, error)
}

func (s *Client) securityLogin(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.Login(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"Login\"")
	}
	req.Header.Set("Login", t.APIKey)
	return nil
}
func (s *Client) securitySecret(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.Secret(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"Secret\"")
	}
	req.Header.Set("Secret", t.APIKey)
	return nil
}
