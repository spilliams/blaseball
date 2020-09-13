package internal

import (
	"github.com/spilliams/blaseball/pkg"
	"github.com/spilliams/blaseball/pkg/model"
)

type LocalDataSession interface {
	pkg.DataSessionDivisionsRead
	dataSessionDivisionsWrite
}

type RemoteDataSession interface {
	pkg.DataSessionDivisionsRead
}

type dataSessionDivisionsWrite interface {
	PutDivision(*model.Division) error
}
