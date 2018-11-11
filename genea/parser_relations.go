package genea

import (
	"fmt"
	"time"

	"github.com/LouisBrunner/regenea/genea/json"
)

func addUnionToPartner(union *Union, partner *Person) {
	if partner == nil {
		return
	}

	partner.Partners = append(partner.Partners, union)
	for _, child := range union.Issues {
		partner.Children = append(partner.Children, child)
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

	// TODO: don't assume that
	child.Father = union.Person1
	child.Mother = union.Person2

	union.Issues = append(union.Issues, child)
	return nil
}

func createUnion(tree *Tree, relation *json.Relation, i int) (*Union, error) {
	union := Union{
		Comments: relation.Comments,
	}

	switch relation.Type {
	case "wedding":
		union.Kind = UnionWedding
	case "civil":
		union.Kind = UnionCivil
	}

	if relation.Begin != nil {
		union.Begin = time.Time(*relation.Begin)
	}
	if relation.End != nil {
		union.End = time.Time(*relation.End)
	}

	person1, err := getPerson(tree, relation.Person1)
	if err != nil {
		return nil, fmt.Errorf("%v (in relation %d)", err, i)
	}
	union.Person1 = person1

	person2, err := getPerson(tree, relation.Person2)
	if err != nil {
		return nil, fmt.Errorf("%v (in relation %d)", err, i)
	}
	union.Person2 = person2

	if union.Person1 != nil && union.Person1 == union.Person2 {
		return nil, fmt.Errorf("relation %d: you can't have an union with yourself", i)
	}

	addUnionToPartner(&union, union.Person1)
	addUnionToPartner(&union, union.Person2)

	return &union, nil
}

func addSibling(child1 *Person, child2 *Person) {
	if !listHasPerson(child1.Siblings, child2) {
		child1.Siblings = append(child1.Siblings, child2)
	}

	if !listHasPerson(child2.Siblings, child1) {
		child2.Siblings = append(child2.Siblings, child1)
	}
}

func addRelations(tree *Tree, relations []json.Relation) error {
	for i, relation := range relations {
		union, err := createUnion(tree, &relation, i)
		if err != nil {
			return err
		}

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
