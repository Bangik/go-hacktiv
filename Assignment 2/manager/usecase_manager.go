package manager

import "hacktiv-assignment-2/usecase"

type UseCaseManager interface {
	OrderUsecase() usecase.OrderUsecase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) OrderUsecase() usecase.OrderUsecase {
	return usecase.NewOrderUsecase(u.repoManager.OrderRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
