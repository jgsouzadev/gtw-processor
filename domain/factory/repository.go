package factory

import "github.com/jgsouzadev/gtw-processor/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
