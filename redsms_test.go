package redsms

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/wildwind123/redsms/ogencl"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("cant load dot env")
	}
}

func TestPing(t *testing.T) {
	t.Skip("manual test")
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	// later ...

	cl, err := ogencl.NewClient("https://cp.redsms.ru", NewSecurity(os.Getenv("REDSMS_LOGIN"), os.Getenv("REDSMS_TOKEN")), ogencl.WithClient(&http.Client{
		Transport: &CustomRoundTripper{Proxied: http.DefaultTransport, Logger: logger},
	}))
	if err != nil {
		t.Error("cant get NewClient", err)
	}
	ctx := context.Background()

	res, err := cl.APIMessagePost(ctx, &ogencl.Request{
		Route: ogencl.RoutePing,
		To:    "+79993332211",
	}, ogencl.APIMessagePostParams{
		Ts: "ts-value",
	})

	responseError := &ogencl.ErrorStatusCode{}

	if errors.As(err, &responseError) {
		t.Error("responseError", responseError)
		return
	}
	if err != nil {
		t.Error("cant send APIMessagePost", err)
		return
	}
	fmt.Println(res)
}

func TestSendSms(t *testing.T) {
	t.Skip("manual test")
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	// later ...

	cl, err := ogencl.NewClient("https://cp.redsms.ru", NewSecurity(os.Getenv("REDSMS_LOGIN"), os.Getenv("REDSMS_TOKEN")), ogencl.WithClient(&http.Client{
		Transport: &CustomRoundTripper{Proxied: http.DefaultTransport, Logger: logger},
	}))
	if err != nil {
		t.Error("cant get NewClient", err)
	}
	ctx := context.Background()

	res, err := cl.APIMessagePost(ctx, &ogencl.Request{
		Route: ogencl.RouteSMS,
		To:    os.Getenv("REDSMS_PHONE"),
		Text:  ogencl.NewOptString("Отправьте код 7777 для входа"),
	}, ogencl.APIMessagePostParams{
		Ts: "ts-value",
	})

	responseError := &ogencl.ErrorStatusCode{}

	if errors.As(err, &responseError) {
		t.Error("responseError", responseError)
		return
	}
	if err != nil {
		t.Error("cant send APIMessagePost", err)
		return
	}
	fmt.Println(res)
}

func TestMessageStatus(t *testing.T) {
	t.Skip("manual test")
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	// later ...

	cl, err := ogencl.NewClient("https://cp.redsms.ru", NewSecurity(os.Getenv("REDSMS_LOGIN"), os.Getenv("REDSMS_TOKEN")), ogencl.WithClient(&http.Client{
		Transport: &CustomRoundTripper{Proxied: http.DefaultTransport, Logger: logger},
	}))
	if err != nil {
		t.Error("cant get NewClient", err)
	}
	ctx := context.Background()

	res, err := cl.APIMessageUUIDGet(ctx, ogencl.APIMessageUUIDGetParams{
		UUID: "c489b8fc-7351-11ef-9a55-0242c0a86496",
	})

	responseError := &ogencl.ErrorStatusCode{}

	if errors.As(err, &responseError) {
		t.Error("responseError", responseError)
		return
	}
	if err != nil {
		t.Error("cant send APIMessagePost", err)
		return
	}
	fmt.Println(res)
}

func TestClientInfo(t *testing.T) {
	t.Skip("manual test")
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	// later ...

	cl, err := ogencl.NewClient("https://cp.redsms.ru", NewSecurity(os.Getenv("REDSMS_LOGIN"), os.Getenv("REDSMS_TOKEN")), ogencl.WithClient(&http.Client{
		Transport: &CustomRoundTripper{Proxied: http.DefaultTransport, Logger: logger},
	}))
	if err != nil {
		t.Error("cant get NewClient", err)
	}
	ctx := context.Background()

	res, err := cl.APIClientInfoGet(ctx, ogencl.APIClientInfoGetParams{
		Ts: "tes",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
}
