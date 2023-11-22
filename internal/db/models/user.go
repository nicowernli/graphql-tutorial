package models

import (
	"github.com/google/uuid"
	graphmodels "github.com/nicowernli/graphql-tutorial/internal/graph/models"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
	RemovedAt *int64    `json:"-"`
}

type UserQuery struct {
	OrderBy *string `json:"orderBy"`
	Sort    *string `json:"sort"`
}

func (u *User) ToGraphqlModel() *graphmodels.User {
	graphUser := &graphmodels.User{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		RemovedAt: u.RemovedAt,
	}

	if u.FirstName == nil {
		graphUser.FirstName = ""
	} else {
		graphUser.FirstName = *u.FirstName
	}

	if u.LastName == nil {
		graphUser.LastName = ""
	} else {
		graphUser.LastName = *u.LastName
	}

	return graphUser
}
