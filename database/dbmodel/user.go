package dbmodel

import (
	"benevolix/pkg/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntry struct {
	gorm.Model `swaggerignore:"true"` // Ignore gorm.Model pour Swagger
	LastName    string            `json:"last_name"`
	FirstName   string            `json:"first_name"`
	Phone       string            `json:"phone" gorm:"uniqueIndex"`
	Email       string            `json:"email" gorm:"uniqueIndex"`
	Password    string            `json:"password"`
	City        string            `json:"city"`
	Bio         string            `json:"bio"`
	Tags        []*TagEntry       `gorm:"many2many:user_tags"`
	Candidature *CandidatureEntry `gorm:"foreignkey:UserID;references:ID"`
}

func (user *UserEntry) ToModel() *model.UserResponse {
	var tags []model.TagResponse
	for _, tag := range user.Tags {
		tags = append(tags, *tag.ToModel())
	}
	return &model.UserResponse{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		City:      user.City,
		Bio:       user.Bio,
		Tags:      tags,
	}
}

type UserRepository interface {
	Create(entry *UserEntry) (*UserEntry, error)
	GetAll() ([]*UserEntry, error)
	GetById(id uint) (*UserEntry, error)
	Update(entry *UserEntry) (*UserEntry, error)
	Delete(id int) error
	GetUserByEmail(email string) (*UserEntry, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(entry *UserEntry) (*UserEntry, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(entry.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	entry.Password = string(hashedPassword)
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *userRepository) GetAll() ([]*UserEntry, error) {
	var entries []*UserEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *userRepository) GetById(id uint) (*UserEntry, error) {
	var entrie *UserEntry
	if err := r.db.First(&entrie, id).Error; err != nil {
		return nil, err
	}
	return entrie, nil
}

func (r *userRepository) Update(entry *UserEntry) (*UserEntry, error) {
	if err := r.db.Save(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&UserEntry{}, id).Error
}

func (r *userRepository) GetUserByEmail(email string) (*UserEntry, error) {
	var entries []*UserEntry
	if err := r.db.Raw("SELECT * FROM user_entries WHERE email = ?;", email).Scan(&entries).Error; err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, errors.New("no email found")
	}
	return entries[0], nil
}
