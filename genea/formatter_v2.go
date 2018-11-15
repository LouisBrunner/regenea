package genea

import (
	"github.com/LouisBrunner/regenea/genea/json"
)

func serializeEventCore(event *EventCore) *json.EventCommon {
	date := json.Date(event.Date)
	return &json.EventCommon{
		Date:     &date,
		Location: event.Location,
		Source:   event.Source,
	}
}

func serializeV2People(tree *Tree) *[]json.PersonV2 {
	people := make([]json.PersonV2, len(tree.People))
	for i, person := range tree.People {
		newPerson := json.PersonV2{
			PersonCommon: serializePersonCommon(person),
			Names:        &person.Names,
		}
		if person.Father != nil {
			newPerson.Parents = &json.Parents{
				Father: json.PersonID(person.Father.ID),
			}
		}
		if person.Mother != nil {
			if newPerson.Parents == nil {
				newPerson.Parents = &json.Parents{}
			}
			newPerson.Parents.Mother = json.PersonID(person.Mother.ID)
		}
		if !person.Birth.Date.IsZero() {
			newPerson.Birth = serializeEventCore(&person.Birth)
		}
		if person.Death != nil {
			newPerson.Death = serializeEventCore(person.Death)
		}
		if len(person.Events) > 0 {
			events := make([]json.Event, len(person.Events))
			for i, event := range person.Events {
				events[i] = json.Event{
					EventCommon: *serializeEventCore(&event.EventCore),
					Title:       event.Title,
				}
			}
			newPerson.Events = &events
		}
		people[i] = newPerson
	}
	return &people
}
func serializeV2Relation(union *Union) json.RelationV2 {
	relation := json.RelationV2{
		RelationCommon: serializeRelationCommon(union),
	}
	if !union.Begin.Date.IsZero() {
		relation.Begin = serializeEventCore(&union.Begin)
	}
	if union.End != nil {
		relation.End = serializeEventCore(union.End)
	}
	return relation
}

func serializeV2Relations(tree *Tree) *[]json.RelationV2 {
	relations := []json.RelationV2{}
	tree.IterateOverUnion(func(union *Union) {
		relations = append(relations, serializeV2Relation(union))
	})
	return &relations
}

func serializeV2(tree *Tree) *json.V2 {
	version := int(VersionV2)
	content := json.V2{}
	content.Version = &version
	content.Comments = tree.Comments
	content.People = serializeV2People(tree)
	content.Relations = serializeV2Relations(tree)
	return &content
}
