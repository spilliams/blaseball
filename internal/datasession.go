package internal

import (
	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

// LocalDataSession represents a session with a local data source (such as a
// database or in-memory store)
type LocalDataSession interface {
	pkg.DataSessionDivisionsRead
	dataSessionDivisionsWrite
}

type dataSessionDivisionsWrite interface {
	PutDivision(*model.Division) error
}
