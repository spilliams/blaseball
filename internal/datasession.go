package internal

import "github.com/spilliams/blaseball/pkg/model"

type LocalDataSession interface {
	DataSessionDivisionsRead
	DataSessionDivisionsWrite
}

type RemoteDataSession interface {
	DataSessionDivisionsRead
}

type DataSessionDivisionsRead interface {
	GetAllDivisions() ([]*model.Division, error)
	GetDivisionByID(string) (*model.Division, error)
}

type DataSessionDivisionsWrite interface {
	PutDivision(*model.Division) error
}
