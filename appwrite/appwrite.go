package appwrite

import (
	"time"

	"github.com/Aloero/Appwrite/client"
	"github.com/Aloero/Appwrite/account"
	"github.com/Aloero/Appwrite/avatars"
	"github.com/Aloero/Appwrite/databases"
	"github.com/Aloero/Appwrite/functions"
	"github.com/Aloero/Appwrite/graphql"
	"github.com/Aloero/Appwrite/health"
	"github.com/Aloero/Appwrite/locale"
	"github.com/Aloero/Appwrite/messaging"
	"github.com/Aloero/Appwrite/storage"
	"github.com/Aloero/Appwrite/teams"
	"github.com/Aloero/Appwrite/users"
)

func NewAccount(clt client.Client) *account.Account {
	return account.New(clt)
}
func NewAvatars(clt client.Client) *avatars.Avatars {
	return avatars.New(clt)
}
func NewDatabases(clt client.Client) *databases.Databases {
	return databases.New(clt)
}
func NewFunctions(clt client.Client) *functions.Functions {
	return functions.New(clt)
}
func NewGraphql(clt client.Client) *graphql.Graphql {
	return graphql.New(clt)
}
func NewHealth(clt client.Client) *health.Health {
	return health.New(clt)
}
func NewLocale(clt client.Client) *locale.Locale {
	return locale.New(clt)
}
func NewMessaging(clt client.Client) *messaging.Messaging {
	return messaging.New(clt)
}
func NewStorage(clt client.Client) *storage.Storage {
	return storage.New(clt)
}
func NewTeams(clt client.Client) *teams.Teams {
	return teams.New(clt)
}
func NewUsers(clt client.Client) *users.Users {
	return users.New(clt)
}

// NewClient initializes a new Appwrite client with a given timeout
func NewClient(optionalSetters ...client.ClientOption) client.Client {
	return client.New(optionalSetters...)
}

// Helper method to construct NewClient()
func WithEndpoint(endpoint string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Endpoint = endpoint
		return nil
	}
}

// Helper method to construct NewClient()
func WithTimeout(timeout time.Duration) client.ClientOption {
	return func(clt *client.Client) error {
		httpClient, err := client.GetDefaultClient(timeout)
		if err != nil {
			return err
		}

		clt.Timeout = timeout
		clt.Client = httpClient

		return nil
	}
}

// Helper method to construct NewClient()
func WithSelfSigned(status bool) client.ClientOption {
	return func(clt *client.Client) error {
		clt.SelfSigned = status
		return nil
	}
}

// Helper method to construct NewClient()
func WithChunkSize(size int64) client.ClientOption {
	return func(clt *client.Client) error {
		clt.ChunkSize = size
		return nil
	}
}

// Helper method to construct NewClient()
// 
// Your project ID
func WithProject(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Project"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your secret API key
func WithKey(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Key"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your secret JSON Web Token
func WithJWT(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-JWT"] = value
		return nil
	}
}
// Helper method to construct NewClient()
func WithLocale(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Locale"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The user session to authenticate with
func WithSession(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Session"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The user agent string of the client that made the request
func WithForwardedUserAgent(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Forwarded-User-Agent"] = value
		return nil
	}
}
