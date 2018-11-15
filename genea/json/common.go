package json

import (
	"time"
)

type Header struct {
	Version *int `json:"version" validate:"required,gt=0"`
}

type PersonID string

type PersonCommon struct {
	ID       PersonID `json:"id" validate:"required"`
	Sex      string   `json:"sex" validate:"required,oneof=M F"`
	Comments string   `json:"comments,omitempty"`
}

type RelationCommon struct {
	Type     string    `json:"type" validate:"required,oneof=wedding civil"`
	Comments string    `json:"comments,omitempty"`
	Person1  *PersonID `json:"person1,omitempty"`
	Person2  *PersonID `json:"person2,omitempty"`
}

type Date time.Time
