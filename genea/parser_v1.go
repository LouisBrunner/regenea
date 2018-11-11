package genea

import (
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

func parseEventV1(dt *json.Date) *EventCore {
	return &EventCore{
		Date: time.Time(*dt),
	}
}

func ImportV1(jsonRepr *json.V1) (*Tree, error) {
	tree := initTree(jsonRepr.Comments, len(*jsonRepr.People))
	for i, jsonPerson := range *jsonRepr.People {
		person := createPerson(&jsonPerson.PersonCommon)
		person.Names.First = jsonPerson.FirstName
		person.Names.Middle = jsonPerson.MiddleName
		person.Names.Last = jsonPerson.LastName
		if jsonPerson.Birthday != nil {
			person.Birth = *parseEventV1(jsonPerson.Birthday)
		}
		if jsonPerson.Deathday != nil {
			person.Death = parseEventV1(jsonPerson.Deathday)
		} else if jsonPerson.Alive != nil && !*jsonPerson.Alive {
			person.Death = &EventCore{}
		}
		err := addPerson(tree, i, person)
		if err != nil {
			return nil, err
		}
	}
	return tree, addRelations(tree, *jsonRepr.Relations)
}
