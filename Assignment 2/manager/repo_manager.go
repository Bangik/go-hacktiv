package manager

import "hacktiv-assignment-2/repository"

type RepoManager interface {
	OrderRepo() repository.OrderRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) OrderRepo() repository.OrderRepository {
	return repository.NewOrderRepository(r.infra.Connection())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
