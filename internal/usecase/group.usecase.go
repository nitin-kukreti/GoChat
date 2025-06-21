package usecase

import "github.com/nitin-kukreti/GoChat/internal/domain"

type GroupUseCase struct {
	repo domain.GroupRepository
}

func NewGroupUseCase(repo domain.GroupRepository) *GroupUseCase {
	return &GroupUseCase{repo: repo}
}

func (uc *GroupUseCase) CreateGroup(name string) (domain.Group, error) {
	return uc.repo.CreateGroup(name)
}

func (uc *GroupUseCase) AddUserToGroup(userID, groupID int) error {
	return uc.repo.AddUserToGroup(userID, groupID)
}
