package config

import "context"

type MockSource struct {
	ReturnMap map[string]string
}

func (m MockSource) Load(_ context.Context) map[string]string {
	return m.ReturnMap
}
