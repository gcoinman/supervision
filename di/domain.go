package di

import (
	"github.com/D-Technologies/go-tokentracker/domain/blocknumber"
	"github.com/D-Technologies/go-tokentracker/infrastructure/db/mysql/blocknumber"
)

var blockNumRepository blocknumberdomain.BlockNumRepository

// InjectBlockNumRepository injects a blocknum repository
func InjectBlockNumRepository() blocknumberdomain.BlockNumRepository {
	if blockNumRepository != nil {
		return blockNumRepository
	}

	blockNumRepository = blocknumber.NewRepository()

	return blockNumRepository
}
