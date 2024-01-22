package pg

import (
	"fmt"
	"net/url"
	"strings"
)

type Option func(*Postgres)

func WithHost(host string) Option {
	return func(postgres *Postgres) {
		postgres.host = host
	}
}

func WithPort(port string) Option {
	return func(postgres *Postgres) {
		postgres.port = port
	}
}

func WithUsername(username string) Option {
	return func(postgres *Postgres) {
		postgres.username = username
	}
}

func WithPassword(password string) Option {
	return func(postgres *Postgres) {
		postgres.password = password
	}
}

func WithDBName(dbName string) Option {
	return func(postgres *Postgres) {
		postgres.dbName = dbName
	}
}

func WithConnectionURI(connectionURI string) Option {
	return func(postgres *Postgres) {
		uri, err := url.Parse(connectionURI)
		if err != nil {
			fmt.Println("Error parsing connection URI:", err)
			return
		}

		postgres.username = uri.User.Username()
		postgres.password, _ = uri.User.Password()
		fmt.Println(postgres.password)
		postgres.host = uri.Hostname()
		postgres.port = uri.Port()
		postgres.dbName = strings.TrimLeft(uri.Path, "/")
	}
}
