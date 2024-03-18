package models

type Country struct {
  Name: `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
}
