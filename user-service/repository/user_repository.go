package repository

import (
	"context"
	"micro-warehouse/user-service/interfaces"
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

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &userRepository{db: db}
}

func (r *userRepository) Createuser(ctx context.Context, user model.User) (*model.User, error) {
	panic("unimplemented")
}

func (r *userRepository) GetAllUser(cxt context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]model.User, int64, error) {
	panic("unimplemented")
}

func (r *userRepository) GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	panic("unimplemented")
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic("unimplemented")
}

func (r *userRepository) UpdateUser(ctx context.Context, user model.User) error {
	panic("unimplemented")
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

func (r *userRepository) GetUserByRole(ctx context.Context, roleName string) ([]model.User, error) {
	panic("unimplemented")
}

func (r *userRepository) AssignUserToRole(ctx context.Context, user_id uuid.UUID, role_id uuid.UUID) error {
	panic("unimplemented")
}

func (r *userRepository) EditAssignToRole(ctx context.Context, assignRoleId uuid.UUID, user_id uuid.UUID, role_id uuid.UUID) error {
	panic("unimplemented")
}

func (r *userRepository) DeleteAssignToRole(ctx context.Context, assignRoleId uuid.UUID) error {
	panic("unimplemented")
}

func (r *userRepository) GetAllUserToRoles(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]model.UserRole, int64, error) {
	panic("unimplemented")
}

func (r *userRepository) GetUserByRoleId(ctx context.Context, roleID uuid.UUID) ([]model.User, error) {
	panic("unimplemented")
}
