package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        string         `gorm:"column:id;primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;default:null"`
}

type ProjectRepository interface {
	Create(ctx context.Context, p *Project) error
	Update(ctx context.Context, p *Project) error
	First(ctx context.Context, id string) (Project, error)
	Find(ctx context.Context) ([]Project, error)
	Delete(ctx context.Context, id string) error
}

type ProjectUsecase interface {
	Create(ctx context.Context, p *Project) error
	Update(ctx context.Context, p *Project) error
	First(ctx context.Context, id string) (Project, error)
	Find(ctx context.Context) ([]Project, error)
	Delete(ctx context.Context, id string) error
}
