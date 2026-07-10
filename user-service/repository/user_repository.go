package repository

import (
	"context"
	"errors"
	"micro-warehouse/user-service/interfaces"
	"micro-warehouse/user-service/model"

	"github.com/gofiber/fiber/v2/log"
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

func (r *userRepository) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return nil, err
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		log.Errorf("{UserRepository} Create User -2 : %w", err)
		return nil, err
	}
	if user.ID == uuid.Nil {
		log.Errorf("{UserRepository} Create User -3 : %w", "user not found")
		return nil, errors.New("user not found")
	}

	return &user, nil
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
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return err
	}
	return nil
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

func checkContext(ctx context.Context, logPrefix string) error {
	select {
	case <-ctx.Done():
		log.Errorf("%s : %w", logPrefix, ctx.Err())
		return ctx.Err()
	default:
		return nil
	}
}
