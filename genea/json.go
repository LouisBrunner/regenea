package genea

import (
	"time"
)

type jsonV1 struct {
	Header
	Comments *string `json:"comments" validate:"omitempty,dive,required"`

	People    *[]jsonPersonV1 `json:"people" validate:"required,dive,min=1"`
	Relations *[]jsonRelation `json:"relations" validate:"required,dive,min=1"`
}

type jsonV2 struct {
	Header
	Comments *string `json:"comments" validate:"omitempty,dive,required"`

	People    *[]jsonPersonV2 `json:"people" validate:"required,dive,min=1"`
	Relations *[]jsonRelation `json:"relations" validate:"required,dive,min=1"`
}

type jsonPersonID string

type jsonPersonCommon struct {
	ID       jsonPersonID `json:"id" validate:"required"`
	Sex      string       `json:"sex" validate:"required,oneof=M F"`
	Comments *string      `json:"comments" validate:"omitempty,dive,required"`
}

type jsonPersonV1 struct {
	jsonPersonCommon
	FirstName  string     `json:"first_name" validate:"required"`
	MiddleName *string    `json:"middle_name" validate:"omitempty,dive,required"`
	LastName   string     `json:"last_name" validate:"required"`
	Birthday   *time.Time `json:"birthday" validate:"omitempty,dive,required"`
	Deathday   *time.Time `json:"deathday" validate:"omitempty,dive,required,gtefield=Birthday"`
	Alive      *bool      `json:"alive" validate:"omitempty"`
}

type jsonPersonV2 struct {
	jsonPersonCommon
	Names  *jsonNames   `json:"names" validate:"required"`
	Birth  *jsonEvent   `json:"birth" validate:"omitempty,dive,required"`
	Death  *jsonEvent   `json:"death" validate:"omitempty,dive,required"`
	Events *[]jsonEvent `json:"events" validate:"omitempty,dive,min=1"`
}

type jsonNames struct {
	First           string  `json:"first" validate:"required"`
	Middle          *string `json:"middle" validate:"omitempty,dive,required"`
	Last            string  `json:"last" validate:"required"`
	AlternativeName *string `json:"alternative" validate:"omitempty,dive,required"`
}

type jsonEvent struct {
	Date     *time.Time `json:"date" validate:"omitempty,dive,required"`
	Location *string    `json:"location" validate:"omitempty,dive,required"`
}

type jsonRelation struct {
	Type     *string         `json:"type" validate:"required,dive,required,oneof=wedding pacs"`
	Person1  *jsonPersonID   `json:"person1" validate:"omitempty,dive,required"`
	Person2  *jsonPersonID   `json:"person2" validate:"omitempty,dive,required"`
	Issues   *[]jsonPersonID `json:"issues" validate:"omitempty,dive,min=1"`
	Begin    *time.Time      `json:"begin" validate:"omitempty,dive,required"`
	End      *time.Time      `json:"end" validate:"omitempty,dive,required"`
	Comments *string         `json:"comments" validate:"omitempty,dive,required"`
}
