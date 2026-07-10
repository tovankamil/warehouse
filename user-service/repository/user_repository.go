package repository

import (
	"context"
	"micro-warehouse/user-service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
feature:
USER ROLE : GET BY EMAIL, GET BY ID, CREATE, UPDATE, DELETE,  GET BY ROLE NAME, GET ALL
ROLE : GET BY ROLE NAME, GET BY ID, CREATE, UPDATE, DELETE,  GET ALL
USER ROLE : ASSIGN USER TO ROLE ,GET USER BY ROLE ID, GET ALL USER ROLE, EDIT ASSIGN TO ROEL

*/

type userRepositoryInterface interface {
	Createuser(ctx context.Context, user model.User) (*model.User, error)
	GetAllUser(cxt context.Context, page int, limit int, search, sortBy, sortOrder string) ([]model.User, int64, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error

	GetUserByRole(ctx context.Context, roleName string) ([]model.User, error)

	AssignUserToRole(ctx context.Context, user_id uuid.UUID, role_id uuid.UUID) error
	EditAssignToRole(ctx context.Context, assignRoleId uuid.UUID, user_id uuid.UUID, role_id uuid.UUID) error
	DeleteAssignToRole(ctx context.Context, assignRoleId uuid.UUID) error
	GetAllUserToRoles(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.UserRole, int64, error)
	GetUserByRoleId(ctx context.Context, roleID uuid.UUID) ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}
