package json

type V1 struct {
	Header
	Comments  string        `json:"comments,omitempty"`
	People    *[]PersonV1   `json:"people" validate:"required,dive,min=1"`
	Relations *[]RelationV1 `json:"relations" validate:"required,dive,min=1"`
}

type PersonV1 struct {
	PersonCommon
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Birthday   *Date  `json:"birthday,omitempty"`
	Deathday   *Date  `json:"deathday,omitempty" validate:"omitempty,gtefield=Birthday"`
	Alive      *bool  `json:"alive,omitempty"`
}

type RelationV1 struct {
	RelationCommon
	Issues *[]PersonID `json:"issues,omitempty" validate:"omitempty,min=1"`
	Begin  *Date       `json:"begin,omitempty"`
	End    *Date       `json:"end,omitempty" validate:"omitempty,gtefield=Begin"`
}
