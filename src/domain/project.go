package domain

import (
	"context"
	"time"
)

type ProjectRequest struct {
	ID          string      `uri:"id"`
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description,omitempty" `
	TechStacks  []TechStack `json:"tech_stacks,omitempty"`
}

type ProjectResponse struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description,omitempty"`
	TechStacks  []TechStack `json:"tech_stacks,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
}

type Project struct {
	ID          string      `firestore:"id"`
	Title       string      `firestore:"title"`
	Description string      `firestore:"description,omitempty"`
	TechStacks  []TechStack `firestore:"tech_stacks,omitempty"`
	UpdatedAt   time.Time   `firestore:"updated_at"`
	CreatedAt   time.Time   `firestore:"created_at"`
}

type ProjectUsecase interface {
	Create(ctx context.Context, p *Project) error
	First(ctx context.Context, id string) (Project, error)
	Find(ctx context.Context) ([]Project, error)
	Update(ctx context.Context, p *Project) error
	Delete(ctx context.Context, id string) error
}

type ProjectRepository interface {
	Create(ctx context.Context, p *Project) error
	First(ctx context.Context, id string) (Project, error)
	Find(ctx context.Context) ([]Project, error)
	Update(ctx context.Context, p *Project) error
	Delete(ctx context.Context, id string) error
}

func (p *Project) BindFromReq(r ProjectRequest) {
	p.Title = r.Title
	p.Description = r.Description
	p.TechStacks = r.TechStacks
}

func (p *Project) BindToRes(r *ProjectResponse) {
	r.ID = p.ID
	r.Title = p.Title
	r.Description = p.Description
	r.CreatedAt = p.CreatedAt
	r.UpdatedAt = p.UpdatedAt
}
