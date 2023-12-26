package usecase

import (
	"context"
	"time"

	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
	customtype "github.com/nawafilhusnul/me-dashboard-api/src/global/custom_type"
)

type projectUsecase struct {
	projectRepo    domain.ProjectRepository
	contextTimeout time.Duration
}

func NewProjectUsecase(a domain.ProjectRepository, timeout time.Duration) domain.ProjectUsecase {
	return &projectUsecase{
		projectRepo:    a,
		contextTimeout: timeout,
	}
}

func (pu *projectUsecase) First(ctx context.Context, id string) (domain.Project, error) {
	var p domain.Project
	return p, nil
}

func (pu *projectUsecase) Find(ctx context.Context) ([]domain.Project, error) {
	pl, err := pu.projectRepo.Find(ctx)
	if err != nil {
		return pl, err
	}

	return pl, nil
}

func (pu *projectUsecase) Update(ctx context.Context, p *domain.Project) error {
	p.UpdatedAt = customtype.SetNullTime(time.Now())
	err := pu.projectRepo.Update(ctx, p)
	return err
}

func (pu *projectUsecase) Delete(ctx context.Context, id string) error {
	err := pu.projectRepo.Delete(ctx, id)
	return err
}

func (pu *projectUsecase) Create(ctx context.Context, p *domain.Project) error {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = customtype.SetNullTime(now)
	err := pu.projectRepo.Create(ctx, p)
	return err
}
