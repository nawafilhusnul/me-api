package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
)

func getUpdateFields(p *domain.Project) []firestore.Update {
	var fields = []firestore.Update{}

	if p.Title != "" {
		fields = append(fields, firestore.Update{Path: "title", Value: p.Title})
	}

	if p.Description != "" {
		fields = append(fields, firestore.Update{Path: "description", Value: p.Description})
	}

	if !p.UpdatedAt.IsZero() {
		fields = append(fields, firestore.Update{Path: "updated_at", Value: p.UpdatedAt})
	}

	return fields
}
