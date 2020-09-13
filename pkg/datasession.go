package pkg

import "github.com/spilliams/blaseball/pkg/model"

type RemoteDataSession interface {
	DataSessionDivisionsRead
}

type DataSessionDivisionsRead interface {
	GetAllDivisions() ([]*model.Division, error)
	GetDivisionByID(string) (*model.Division, error)
	GetDivisionByName(string) (*model.Division, error)
}
