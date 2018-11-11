package json

type V2 struct {
	Header
	Comments  string      `json:"comments,omitempty"`
	People    *[]personV2 `json:"people" validate:"required,dive,min=1"`
	Relations *[]Relation `json:"relations" validate:"required,dive,min=1"`
}

type personV2 struct {
	PersonCommon
	Names  *Names       `json:"names" validate:"required"`
	Birth  *EventCommon `json:"birth,omitempty"`
	Death  *EventCommon `json:"death,omitempty"`
	Events *[]Event     `json:"events,omitempty" validate:"omitempty,dive,min=1"`
}

type Names struct {
	First           string `json:"first,omitempty"`
	Middle          string `json:"middle,omitempty"`
	Last            string `json:"last,omitempty"`
	AlternativeName string `json:"alternative,omitempty"`
}

type EventCommon struct {
	Date     *Date  `json:"date,omitempty"`
	Location string `json:"location,omitempty"`
	Source   string `json:"source,omitempty"`
}

type Event struct {
	EventCommon
	Title string `json:"title" validate:"required"`
}
