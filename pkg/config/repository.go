package config

import (
	"context"
	"strconv"
)

// Repository is a entity that helps us to manage variables returned from a Source.
type Repository struct {
	values map[string]string
}

// NewRepository instantiates a new Repository.
func NewRepository() *Repository {
	return &Repository{
		values: make(map[string]string),
	}
}

// Source will load all the variables returned from a Source in the Repository.
func (r *Repository) Source(ctx context.Context, getter Source) {
	for k, v := range getter.Load(ctx) {
		r.Add(k, v)
	}
}

// Add apply the given key in a parse function and then add it with the value on
// the Repository.
func (r *Repository) Add(k string, v string) {
	r.values[k] = v
}

// Get will seach for the given key in the values map on the Repository.
func (r Repository) Get(key string) string {
	return r.values[key]
}

// GetDefault does the same that Get does, except it returns a default value if
// not found the given key in the values map.
func (r Repository) GetDefault(key, d string) string {
	v, ok := r.values[key]
	if !ok {
		return d
	}

	return v
}

// GetInt will seach for the given key in the values map on the Repository and convert
// its value to an int.
func (r Repository) GetInt(key string) int {
	v, ok := r.values[key]
	if !ok {
		return 0
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}

	return i
}

// GetIntDefault does the same that GetInt does, except it returns a default value
// if not found the given key in the values map.
func (r Repository) GetIntDefault(key string, d int) int {
	v, ok := r.values[key]
	if !ok {
		return d
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return d
	}

	return i
}

// GetBool will seach for the given key in the values map on the Repository and convert
// its value to a bool.
func (r Repository) GetBool(key string) bool {
	v, ok := r.values[key]
	if !ok {
		return false
	}

	if v == "true" {
		return true
	}

	return false
}

// GetBoolDefault does the same that GetBool does, except it returns a default value
// if not found the given key in the values map.
func (r Repository) GetBoolDefault(key string, d bool) bool {
	v, ok := r.values[key]
	if !ok {
		return d
	}

	if v == "true" {
		return true
	}

	if v == "false" {
		return false
	}

	return d
}
