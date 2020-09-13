package memdata

import (
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg/model"
)

type MemoryDataSession struct {
	allDivisions []*model.Division
}

func NewSession() internal.LocalDataSession {
	return &MemoryDataSession{
		allDivisions: []*model.Division{},
	}
}
