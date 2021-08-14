package config

import "context"

// Source can be used to load env variables from a defined source.
type Source interface {
	Load(context.Context) map[string]string
}
