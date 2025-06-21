package domain

type GroupRepository interface {
	CreateGroup(name string) (Group, error)
	AddUserToGroup(userID, groupID int) error
}
