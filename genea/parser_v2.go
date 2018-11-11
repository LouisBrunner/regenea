package genea

import (
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

func parseEventCoreV2(jsonEvent *json.EventCommon) *EventCore {
	event := EventCore{
		Location: jsonEvent.Location,
		Source:   jsonEvent.Source,
	}
	if jsonEvent.Date != nil {
		event.Date = time.Time(*jsonEvent.Date)
	}
	return &event
}

func parseEventV2(jsonEvent *json.Event) *Event {
	return &Event{
		EventCore: *parseEventCoreV2(&jsonEvent.EventCommon),
		Title:     jsonEvent.Title,
	}
}

func ImportV2(jsonRepr *json.V2) (*Tree, error) {
	tree := initTree(jsonRepr.Comments, len(*jsonRepr.People))
	for i, jsonPerson := range *jsonRepr.People {
		person := createPerson(&jsonPerson.PersonCommon)
		person.Names = *jsonPerson.Names
		if jsonPerson.Birth != nil {
			person.Birth = *parseEventCoreV2(jsonPerson.Birth)
		}
		if jsonPerson.Death != nil {
			person.Death = parseEventCoreV2(jsonPerson.Death)
		}
		if jsonPerson.Events != nil {
			person.Events = make([]Event, len(*jsonPerson.Events))
			for j, jsonEvent := range *jsonPerson.Events {
				person.Events[j] = *parseEventV2(&jsonEvent)
			}
		}
		err := addPerson(tree, i, person)
		if err != nil {
			return nil, err
		}
	}
	return tree, addRelations(tree, *jsonRepr.Relations)
}
