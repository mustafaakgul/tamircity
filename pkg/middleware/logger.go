package middleware

import (
	"github.com/getsentry/sentry-go"
	"log"
)

func SentryLogger() {
	err_sentry := sentry.Init(sentry.ClientOptions{
		Dsn: "YOUR_DSN",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err_sentry != nil {
		log.Fatalf("sentry.Init: %s", err_sentry)
	}
	sentry.CaptureMessage("It works!")
}
