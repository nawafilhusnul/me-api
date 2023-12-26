package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	customtype "github.com/nawafilhusnul/me-dashboard-api/src/global/custom_type"
	"gorm.io/gorm"
)

type Project struct {
	ID            uuid.UUID             `gorm:"primarykey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title         string                `gorm:"column:title;not null;type:varchar(255)" json:"title"`
	ThumbnailUrl  customtype.NullString `gorm:"column:thumbnail_url;type:varchar(255);default:null" json:"thumbnail_url"`
	About         string                `gorm:"column:about;not null;type:text" json:"about"`
	Year          string                `gorm:"column:year;not null;type:varchar(4)" json:"year"`
	SourceCodeUrl customtype.NullString `gorm:"column:source_code_url;type:varchar(255);default:null" json:"source_code_url"`
	DeploymentUrl customtype.NullString `gorm:"column:deployment_url;type:varchar(255);default:null" json:"deployment_url"`
	CreatedAt     time.Time             `gorm:"column:created_at;default:current_timestamp" json:"created_at"`
	UpdatedAt     sql.NullTime          `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt        `gorm:"column:deleted_at" json:"deleted_at"`
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
