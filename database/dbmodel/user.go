package dbmodel

import "gorm.io/gorm"

type UserEntry struct {
	gorm.Model
	LastName    string            `json:"last_name"`
	FirstName   string            `json:"first_name"`
	Phone       string            `json:"phone"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	City        string            `json:"city"`
	Bio         string            `json:"bio"`
	Tags        []*TagEntry       `gorm:"many2many:user_tags"`
	Candidature *CandidatureEntry `gorm:"foreignkey:UserID;references:ID"`
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
	}
	return entries[0], nil
}
