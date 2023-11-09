package firestore

import (
	"time"

	"cloud.google.com/go/firestore"
	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
)

var (
	ZeroTime = time.Date(1, 1, 1, 0, 0, 0, 0, &time.Location{})
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
