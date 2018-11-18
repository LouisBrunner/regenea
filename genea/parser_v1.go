package genea

import (
	"fmt"
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

func parseEventV1(dt *json.Date) *EventCore {
	return &EventCore{
		Date: time.Time(*dt),
	}
}

func addFatherOrMother(parent *Person, child *Person) {
	if parent.Sex == SexMale {
		child.Father = parent
	} else if parent.Sex == SexFemale {
		child.Mother = parent
	}
}

func addChildrenToUnion(tree *Tree, union *Union, i int, childID json.PersonID) error {
	child, err := getPerson(tree, &childID)
	if err != nil {
		return fmt.Errorf("%v (in relation %d)", err, i)
	}

	if child.flags.issuesFound {
		return fmt.Errorf("relation %d: `%s` is already listed in another `issues`", i, child.ID)
	}
	child.flags.issuesFound = true

	sex1 := SexOther
	child.Family = union
	if union.Person1 != nil {
		sex1 = union.Person1.Sex
		addFatherOrMother(union.Person1, child)
	}
	if union.Person2 != nil {
		if sex1 == union.Person2.Sex {
			return fmt.Errorf("relation %d: two partners cannot have the same sex", i)
		}
		addFatherOrMother(union.Person2, child)
	}

	union.Issues = append(union.Issues, child)
	return nil
}

func addRelationsV1(tree *Tree, relations []json.RelationV1) error {
	for i, relation := range relations {
		union, err := createUnion(tree, &relation.RelationCommon, i)
		if err != nil {
			return err
		}

		if relation.Begin != nil {
			union.Begin = *parseEventV1(relation.Begin)
		}
		if relation.End != nil {
			union.End = parseEventV1(relation.End)
		}
		// TODO: check dates order

		if relation.Issues != nil {
			for _, childID := range *relation.Issues {
				err = addChildrenToUnion(tree, union, i, childID)
				if err != nil {
					return err
				}
			}
		}

		for i, child := range union.Issues {
			for _, child2 := range union.Issues[i+1:] {
				addSibling(child, child2)
			}
		}
	}
	return nil
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
	return tree, addRelationsV1(tree, *jsonRepr.Relations)
}
