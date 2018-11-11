package json

type V1 struct {
	Header
	Comments  string      `json:"comments,omitempty"`
	People    *[]personV1 `json:"people" validate:"required,dive,min=1"`
	Relations *[]Relation `json:"relations" validate:"required,dive,min=1"`
}

type personV1 struct {
	PersonCommon
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Birthday   *Date  `json:"birthday,omitempty"`
	Deathday   *Date  `json:"deathday,omitempty" validate:"omitempty,gtefield=Birthday"`
	Alive      *bool  `json:"alive,omitempty"`
}
