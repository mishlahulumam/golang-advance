package postgres_gorm

import (
	"context"
	"errors"
	"golang-advance/assignment-3/entity"
	"golang-advance/assignment-3/service"
	"log"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type shortlinkRepository struct {
	db GormDBIface
}

func NewShortlinkRepository(db GormDBIface) service.IShortlinkRepository {
	return &shortlinkRepository{db: db}
}

func (r *shortlinkRepository) CreateShortlink(ctx context.Context, shortlink *entity.Shortlink) (entity.Shortlink, error) {
	if err := r.db.WithContext(ctx).Create(shortlink).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.Shortlink{}, err
	}
	return *shortlink, nil
}

func (r *shortlinkRepository) GetLongURL(ctx context.Context, longURL string) (entity.Shortlink, error) {
	var shortlink entity.Shortlink
	if err := r.db.WithContext(ctx).Where("shortlink = ?", longURL).First(&shortlink).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shortlink, nil
		}
		log.Printf("Error getting wallet by user ID: %v\n", err)
		return entity.Shortlink{}, err
	}
	return shortlink, nil
}
