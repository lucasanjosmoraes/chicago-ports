package database

import (
	"context"
	"database/sql"
)

// Opener defines what is needed to instantiate a new database/sql.DB. it doesn't
// implement methods from sttopage.Stopper interface because they should be defined
// in a Entity Repository and implemented by its Data Sources.
type Opener interface {
	Open(context.Context) (*sql.DB, error)
}
