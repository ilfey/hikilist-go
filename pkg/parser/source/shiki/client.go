package shiki

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/source/shiki/config"
	"github.com/sirupsen/logrus"
)

func newClient(logger logrus.FieldLogger, config *config.Config) *resty.Client {
	client := resty.New()

	client.SetTimeout(config.RequestTimeout * time.Millisecond)

	client.SetContentLength(true)

	client.SetBaseURL(config.BaseUrl)

	client.SetHeaders(map[string]string{
		"User-Agent":   config.UserAgent,
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		logger.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL,
		}).Debug("Started")

		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		logger.WithFields(logrus.Fields{"status": r.Status(), "method": r.Request.Method, "url": r.Request.URL, "time": r.Time()}).Debug("Completed")

		if r.StatusCode() >= 400 && r.StatusCode() < 600 {
			return &extractor.NetworkError{
				StatusCode: r.StatusCode(),
			}
		}

		return nil
	})

	return client
}
