package genea

import (
	"github.com/LouisBrunner/regenea/genea/json"
)

func serializeV1People(tree *Tree) *[]json.PersonV1 {
	people := make([]json.PersonV1, len(tree.People))
	for i, person := range tree.People {
		newPerson := json.PersonV1{
			PersonCommon: serializePersonCommon(person),
			FirstName:    person.Names.First,
			MiddleName:   person.Names.Middle,
			LastName:     person.Names.Last,
		}
		if !person.Birth.Date.IsZero() {
			date := json.Date(person.Birth.Date)
			newPerson.Birthday = &date
		}
		if person.Death != nil {
			if !person.Death.Date.IsZero() {
				date := json.Date(person.Death.Date)
				newPerson.Deathday = &date
			} else {
				False := false
				newPerson.Alive = &False
			}
		}
		people[i] = newPerson
	}
	return &people
}

func serializeV1Relation(union *Union) json.RelationV1 {
	relation := json.RelationV1{
		RelationCommon: serializeRelationCommon(union),
	}
	if !union.Begin.Date.IsZero() {
		date := json.Date(union.Begin.Date)
		relation.Begin = &date
	}
	if union.End != nil {
		date := json.Date(union.End.Date)
		relation.End = &date
	}
	if len(union.Issues) > 0 {
		issues := []json.PersonID{}
		for _, issue := range union.Issues {
			issues = append(issues, json.PersonID(issue.ID))
		}
		relation.Issues = &issues
	}
	return relation
}

func serializeV1Relations(tree *Tree) *[]json.RelationV1 {
	relations := []json.RelationV1{}
	tree.IterateOverUnion(func(union *Union) {
		relations = append(relations, serializeV1Relation(union))
	})
	return &relations
}

func serializeV1(tree *Tree) *json.V1 {
	version := int(VersionV1)
	content := json.V1{}
	content.Version = &version
	content.Comments = tree.Comments
	content.People = serializeV1People(tree)
	content.Relations = serializeV1Relations(tree)
	return &content
}
