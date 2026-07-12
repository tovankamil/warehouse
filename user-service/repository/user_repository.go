package repository

import (
	"context"
	"errors"
	"micro-warehouse/user-service/interfaces"
	"micro-warehouse/user-service/model"
	"time"

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

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if user.ID == uuid.Nil {
			return errors.New("user not found")
		}

		return nil
	})

	if err != nil {
		log.Errorf("{UserRepository} Create User -2 : %w", err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]model.User, int64, error) {
	if err := checkContext(ctx, "{UserRepository} GetAll User -1"); err != nil {
		return nil, 0, err
	}

	if page <= 0 {
		page = 1
	}

	if limit <= 10 {
		limit = 10
	}

	if sortBy == "" {
		sortBy = "created_at"
	}

	if sortOrder == "" {
		sortOrder = "desc"
	}

	//calculate offset

	offset := (page - 1) * limit
	query := r.db.WithContext(ctx).Model(&model.User{})
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	var totalRecords int64
	if err := query.Count(&totalRecords).Error; err != nil {
		log.Errorf("{UserRepository} GetAll User -2 :%v", err)
		return nil, 0, err
	}
	modelUsers := []model.User{}
	if err := query.Select("id", "name", "email", "password", "photo", "phone", "created_at").
		Preload("Roles").
		Offset(offset).
		Limit(limit).
		Find(&modelUsers).Error; err != nil {
		log.Errorf("{Repository} Getalluser - 2 : %v", err)
		return nil, 0, err
	}
	return modelUsers, totalRecords, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return nil, err
	}
	modelUsers := model.User{}
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "photo", "phone", "created_at").
		Where("id = ?", id).
		Preload("Roles").
		First(&modelUsers).Error; err != nil {
		log.Errorf("{Repository} GetUserById - 1 : %v", err)
		return nil, err
	}
	return &modelUsers, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return nil, err
	}
	modelUsers := model.User{}
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "photo", "phone", "created_at").
		Where("email = ?", email).
		Preload("Roles").
		First(&modelUsers).Error; err != nil {
		log.Errorf("{Repository} GetUserByEmail - 1 : %v", err)
		return nil, err
	}
	return &modelUsers, nil

}

func (r *userRepository) UpdateUser(ctx context.Context, user model.User) error {
	if err := checkContext(ctx, "{UserRepository} Update User -1"); err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"name":  user.Name,
		"phone": user.Phone,
		"photo": user.Photo,
	}
	if user.Password != "" {
		updateData["password"] = user.Password
	}

	res := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", user.ID).Updates(updateData)
	if res.Error != nil {
		log.Errorf("{Repository} Update User -2 :%v", res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("data user tidak ditemukan")
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return err
	}
	modelUser := model.User{}

	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "photo", "phone").
		Preload("Roles").Where("id = ?", id).First(&modelUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("data user tidak ditemukan")
		}
		log.Errorf("{Repository} Delete User -2 :%v", err)
		return err
	}

	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

func (r *userRepository) GetUserByRole(ctx context.Context, roleName string) ([]model.User, error) {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return nil, err
	}

	users := []model.User{}

	subQuery := r.db.Table("user_role").
		Select("user_role.user_id").
		Joins("JOIN roles on user_role.role_id = roles.id").
		Where("roles.name = ?", roleName)

	if err := r.db.WithContext(ctx).
		Where("id IN ()", subQuery).
		Preload("Roles").
		Find(&users).Error; err != nil {
		log.Errorf("{Repository} GetUserByRole -2 :%v", err)
		return nil, err
	}
	return users, nil
}

func (r *userRepository) AssignUserToRole(ctx context.Context, user_id uuid.UUID, role_id uuid.UUID) error {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return err
	}

	userRole := model.UserRole{
		UserID:     user_id,
		RoleID:     role_id,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	if err := r.db.WithContext(ctx).Create(&userRole).Error; err != nil {
		log.Errorf("{Repository} AssignUserToRole -2 :%v", err)
		return err
	}

	return nil
}

func (r *userRepository) EditAssignToRole(ctx context.Context, assignRoleId uuid.UUID, user_id uuid.UUID, role_id uuid.UUID) error {
	if err := checkContext(ctx, "{UserRepository} EditAssignToRole -1"); err != nil {
		return err
	}

	res := r.db.WithContext(ctx).Model(&model.UserRole{}).Where("id = ?", assignRoleId).Updates(map[string]interface{}{
		"user_id":    user_id,
		"role_id":    role_id,
		"updated_at": time.Now(),
	})
	if res.Error != nil {
		log.Errorf("{Repository} EditAssignToRole -2 :%v", res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("data user role tidak ditemukan")
	}

	return nil
}

func (r *userRepository) DeleteAssignToRole(ctx context.Context, assignRoleId uuid.UUID) error {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetAllUserToRoles(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]model.UserRole, int64, error) {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return nil, 0, err
	}
	userRoles := []model.UserRole{}
	var totalRecords int64

	query := r.db.WithContext(ctx).Model(&model.UserRole{})
	if search != "" {
		query = query.Joins("JOIN users on users.id = user_role.user_id").
			Joins("JOIN roles on roles.id = user_role.role_id").
			Where("users.name LIKE ?", "%"+search+"%")
	}
	if err := query.Count(&totalRecords).Error; err != nil {
		log.Errorf("{Repository} Getallusertoroles -2 :%v", err)
		return nil, 0, err
	}

	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}
	offset := (page - 1) * limit
	if err := query.Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&userRoles).Error; err != nil {
		log.Errorf("{Repository} Getallusertoroles -3 :%v", err)
		return nil, 0, err
	}

	return userRoles, totalRecords, nil
}

func (r *userRepository) GetUserByRoleId(ctx context.Context, assignRoleId uuid.UUID) (*model.UserRole, error) {
	if err := checkContext(ctx, "{UserRepository} Create User -1"); err != nil {
		return nil, err
	}

	userRole := model.UserRole{}

	if err := r.db.WithContext(ctx).Select("id", "user_id", "role_id", "updated_at").
		Preload("User").
		Preload("Role").
		Where("id = ?", assignRoleId).First(&userRole).Error; err != nil {
		log.Errorf("{Repository} GetUserByRoleId - 2 :%v", err)
		return nil, err
	}
	return &userRole, nil
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
