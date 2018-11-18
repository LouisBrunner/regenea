package json

type V2 struct {
	Header
	Comments  string        `json:"comments,omitempty"`
	People    *[]PersonV2   `json:"people" validate:"required,dive,min=1"`
	Relations *[]RelationV2 `json:"relations" validate:"required,dive,min=1"`
}

type PersonV2 struct {
	PersonCommon
	Names   *Names       `json:"names" validate:"required"`
	Parents *Parents     `json:"parents,omitempty"`
	Birth   *EventCommon `json:"birth,omitempty"`
	Death   *EventCommon `json:"death,omitempty"`
	Events  *[]Event     `json:"events,omitempty" validate:"omitempty,dive,min=1"`
}

type Names struct {
	First           string `json:"first,omitempty"`
	Middle          string `json:"middle,omitempty"`
	Last            string `json:"last,omitempty"`
	Alternative string `json:"alternative,omitempty"`
}

type Parents struct {
	Father PersonID `json:"father,omitempty"`
	Mother PersonID `json:"mother,omitempty"`
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

type RelationV2 struct {
	RelationCommon
	Begin *EventCommon `json:"begin,omitempty"`
	End   *EventCommon `json:"end,omitempty"`
}
