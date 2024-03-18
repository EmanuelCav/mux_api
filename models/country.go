package models

type Country struct {
  Name string `json:"name,omitempty" bson:"name,omitempty"`
}
