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

func addChildrenOf(parent *Person, union *Union) bool {
	if parent == nil {
		return false
	}

	otherParent := union.Person2
	if parent == otherParent {
		otherParent = union.Person1
	}

	for _, child := range parent.Children {
		if child.Father == parent && child.Mother == otherParent ||
			child.Father == otherParent && child.Mother == parent {
			union.Issues = append(union.Issues, child)
		}
	}

	return true
}

func addRelationsV2(tree *Tree, relations []json.RelationV2) error {
	for i, relation := range relations {
		union, err := createUnion(tree, &relation.RelationCommon, i)
		if err != nil {
			return err
		}

		if relation.Begin != nil {
			union.Begin = *parseEventCoreV2(relation.Begin)
		}
		if relation.End != nil {
			union.End = parseEventCoreV2(relation.End)
		}

		if !addChildrenOf(union.Person1, union) {
			addChildrenOf(union.Person2, union)
		}
	}
	return nil
}

func addParentsToChildren(tree *Tree, people []json.PersonV2) error {
	for i, person := range tree.People {
		jsonPerson := people[i]
		if jsonPerson.Parents != nil {
			parent, err := getPerson(tree, &jsonPerson.Parents.Father)
			if err != nil {
				return err
			}
			person.Father = parent
			if parent != nil {
				person.Father.Children = append(person.Father.Children, person)
				for _, child := range person.Father.Children {
					addSibling(person, child)
				}
			}

			parent, err = getPerson(tree, &jsonPerson.Parents.Mother)
			if err != nil {
				return err
			}
			person.Mother = parent
			if parent != nil {
				person.Mother.Children = append(person.Mother.Children, person)
				for _, child := range person.Mother.Children {
					addSibling(person, child)
				}
			}
		}
	}
	return nil
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

	err := addParentsToChildren(tree, *jsonRepr.People)
	if err != nil {
		return nil, err
	}

	return tree, addRelationsV2(tree, *jsonRepr.Relations)
}
