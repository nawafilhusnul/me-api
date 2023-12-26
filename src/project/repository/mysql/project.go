package mysql

import (
	"context"

	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
	"gorm.io/gorm"
)

type projectRepository struct {
	conn *gorm.DB
}

func NewProjectRepository(conn *gorm.DB) domain.ProjectRepository {
	return &projectRepository{conn}
}

func (pr *projectRepository) First(ctx context.Context, id string) (domain.Project, error) {
	var p domain.Project
	// TODO: implementing soon
	return p, nil
}

func (pr *projectRepository) Find(ctx context.Context) ([]domain.Project, error) {
	var p = []domain.Project{}
	err := pr.conn.Find(&p).Error
	if err != nil {
		return p, err
	}

	return p, nil
}

func (pr *projectRepository) Update(ctx context.Context, p *domain.Project) error {
	// TODO: implementing soon
	return nil
}

func (pr *projectRepository) Delete(ctx context.Context, id string) error {
	// TODO: implementing soon
	return nil
}

func (pr *projectRepository) Create(ctx context.Context, p *domain.Project) error {
	// TODO: implementing soon
	return nil
}
