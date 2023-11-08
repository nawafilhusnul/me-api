package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
)

type firestoreProjectRepository struct {
	Conn *firestore.Client
}

func NewFirestoreProjectRepository(conn *firestore.Client) domain.ProjectRepository {
	return &firestoreProjectRepository{conn}
}

func (pu *firestoreProjectRepository) First(ctx context.Context, id string) (domain.Project, error) {
	var p domain.Project
	return p, nil
}

func (pu *firestoreProjectRepository) Find(ctx context.Context) ([]domain.Project, error) {
	var p = []domain.Project{}
	return p, nil
}

func (pu *firestoreProjectRepository) Update(ctx context.Context, p *domain.Project) error {
	if p == nil {
		return domain.ErrBadParamInput
	}

	ref := pu.Conn.Collection("projects").Doc(p.ID)
	err := pu.Conn.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		_, err := t.Get(ref)
		if err != nil {
			return err
		}

		return t.Update(ref, getUpdateFields(p))
	})

	if err != nil {
		return err
	}
	return nil
}

func (pu *firestoreProjectRepository) Delete(ctx context.Context, id string) error {
	if id == "" {
		return domain.ErrBadParamInput
	}

	ref := pu.Conn.Collection("projects").Doc(id)
	err := pu.Conn.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		_, err := t.Get(ref)
		if err != nil {
			return err
		}

		return t.Update(ref, []firestore.Update{
			{
				Path: "deleted_at", Value: time.Now(),
			},
		})
	})

	if err != nil {
		return err
	}
	return nil
}

func (pu *firestoreProjectRepository) Create(ctx context.Context, p *domain.Project) error {
	ref := pu.Conn.Collection("projects").NewDoc()
	p.ID = ref.ID

	_, err := ref.Set(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
