package redsms

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/wildwind123/redsms/ogencl"
)

// security login

type SecurityLogin struct {
	login string
	token string
}

func NewSecurity(login, token string) *SecurityLogin {
	return &SecurityLogin{
		login: login,
		token: token,
	}
}

func (sL *SecurityLogin) Login(ctx context.Context, operationName string) (ogencl.Login, error) {
	return ogencl.Login{
		APIKey: sL.login,
	}, nil
}

func (sL *SecurityLogin) Secret(ctx context.Context, operationName string) (ogencl.Secret, error) {
	return ogencl.Secret{
		APIKey: sL.token,
	}, nil
}

// round  trip

type CustomRoundTripper struct {
	Proxied http.RoundTripper
	Logger  *slog.Logger
}

// RoundTrip executes before and after the HTTP request
func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	secret := req.Header.Get("Secret")
	ts := req.Header.Get("Ts")
	req.Header.Set("Secret", GenerateSecret(ts, secret))

	r, err := c.Proxied.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	// Copy the response body
	var bodyCopy []byte
	if r.Body != nil {
		bodyCopy, err = io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		if c.Logger != nil {
			c.Logger.Debug("response", slog.String("body", string(bodyCopy)))
		}

		// Replace the original response body with a new io.ReadCloser containing the copy
		r.Body = io.NopCloser(bytes.NewBuffer(bodyCopy))
	}

	return r, err
}

// generate secret

func GenerateSecret(ts, secret string) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%s%s", ts, secret)))
	return hex.EncodeToString(hasher.Sum(nil))
}
