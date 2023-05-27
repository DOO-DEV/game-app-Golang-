package entity

// Access control only keeps allowed permission

type AccessControl struct {
	ID           uint
	ActorID      uint
	ActorType    ActorType
	PermissionID uint
}

type ActorType string

const (
	RoleActorType = "role"
	UserActorType = "user"
)
