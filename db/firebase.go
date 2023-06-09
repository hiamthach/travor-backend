package db

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	App     *firebase.App
	Storage *storage.Client
)

func InitializeFirebase() error {
	opt := option.WithCredentialsJSON([]byte(`
		{
			"type": "service_account",
			"project_id": "travor-4704d",
			"private_key_id": "fe748f0f042f251719528cabcaf9ca0ad8c10cd7",
			"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCs9VRhrFSceww7\nE5/IvrlcYJcnpnA6PM6OMh6ZePT6Fzx9+UvAUh8nuy0SYAzUMgoAIC4v0o+PEU55\nJeTf1i8C/if01wMEPBUbBCi/jp7FVdNQs0rTxWxUY6iwlcFsFL3DCMzx/ghlE8HU\nkRM04tZZYIOoSgzLY1y4sbiiYtou5cV+jW1WxFGQdmWJsvHjRYZM4vBGbjO7WSST\nCeth8olOTtPxmVw4gkKv/jJWVpjUpLLPouv+Po/jEHMqZiEgM94DMNfwNltaQvxi\nDS9AffDi3nTjT5r59vgfcKCFK+Z9K19nRoCoFT3/+1tuKyxH59258rJYOmezLUOm\nFzyAEFcrAgMBAAECggEATqfTFAg3AmLDESq/g5+Y1HXox9NdPl7g4LHSUKE9Z26h\nx/hThsbR/FOXD8A2lVmcSqpX1s+/EaUhNt8Q0uqovoeuzB3r9UUNpBekIFlPwxZg\n31ZqNRyXQ1l/Ia85I1nbYpLiATsxviXaBD1lqqtuJ39I4IOsheJODIBmIxMpfiGo\nb+WEPqyGyOwxp+wo+7AcJCwKhJdvRi+dG7+r4X5adaVYeJuJ6mj7j1yVTQ0Oq28m\nhYD4P+LijT5/DT5IcX12XPkR55MONV/DOSIYDSBtoTZGbX627OgwBia70Gd5yMwP\nZHCa03SEiUg8/BNv0LqN1aHvGMJYi/BhCcqD+5CTYQKBgQDhp3c6B2WVLPd5KRC9\nhv0K9zNc5YHhSaMXSKykK05J6WILKlhVrTaHfUU+aLhH3PEY+SMo+lzl7oVVnTQA\nBdq0pWbaFiDLBu4wIZc3eXF3w1kKM6RF4qacovFkitzwazWZrLhzFq5iVvyzuHKy\neaRjNnYLSngAPMwEBf0zWyqhjQKBgQDEN7fr4OcxS7qddOOdL/yNcFiPAn25tbb8\nUr2wecq6kSUlsw0xO/BSk16no8wvtHtovCaoLKMlZtfD4UIFwY0lbbxlQ2rBfOMj\nGcM6eQSdqm46cyfF0oC5pBZ+rHWysTFQFVQHatuj6pO/nsQW9PboSQMN1EQyXCgr\n/Qsb9E+BlwKBgQC1qTP6q2nEJmc0SzMG4V1lLo1TBBEDZVy+qesSXyqvpgsgIf7r\nOn7jpJ3SGEwCzoPqLud3Xdbb5KHCwsPSIORo+Z9LgvTzfajTZrduGX6Xg9RyuhBN\nkqG/CA59eq1Tf3uvM7VqMaBB2vWlWaA8vmDw3i8598GTr14u/wNQDNzNIQKBgBYP\nTZYvSs63/NsG8TAS3Qnpv4sUdJ4UV+Y2Ry04BBrFhq+eWbVXT8V43ANs3t/eIqG7\nxXzSx+TEky1qfcm34O58T9nyomXubHrdwp0oTeSL7KiVmbsb120HjZ/gM7tLQIgI\nXdIXqWncTZAVgWYzVC75RmafnqsqwTlw0XbrGKUxAoGAD5C8GVMH89yEfBv5CJt2\n9IGADYRTpVr+l/ftgv1gQyuS0Yu8ZKzwKk3i0Ta/PoTMobfO+5Y0BEk5Tmdk6gFp\nLkskv35pzmEG+nM8eSN7d/zBx2N9q6FSPbR9TtZyC0huvi1TXz8T54QIVwj61nS2\nKAzA65Q9O+e370W2iaCobYI=\n-----END PRIVATE KEY-----\n",
			"client_email": "firebase-adminsdk-8ewsh@travor-4704d.iam.gserviceaccount.com",
			"client_id": "107084355128647854035",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-8ewsh%40travor-4704d.iam.gserviceaccount.com",
			"universe_domain": "googleapis.com"
		}
	`))

	// Initialize Firebase Admin SDK.
	var err error
	App, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	// Initialize Firebase Storage client.
	Storage, err = App.Storage(context.Background())
	if err != nil {
		return err
	}

	return nil
}
