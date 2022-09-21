package middleware

import (
	"github.com/getsentry/sentry-go"
	"log"
)

func SentryLogger() {
	err_sentry := sentry.Init(sentry.ClientOptions{
		Dsn: "https://c23b897606e945f1aaa3f0440f523d34@o1418275.ingest.sentry.io/6766496",
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
