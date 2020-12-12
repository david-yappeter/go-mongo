// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Name string `json:"name"`
}

type UpdateUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserOps struct {
	Create      *User   `json:"create"`
	CreateBatch []*User `json:"create_batch"`
	Update      *User   `json:"update"`
	Delete      string  `json:"delete"`
}