package genea

import (
	"fmt"

	"github.com/LouisBrunner/regenea/genea/json"
)

func addUnionToPartner(union *Union, partner *Person) {
	if partner == nil {
		return
	}

	partner.Partners = append(partner.Partners, union)
	for _, child := range union.Issues {
		if !listHasPerson(partner.Children, child) {
			partner.Children = append(partner.Children, child)
		}
	}
}

func createUnion(tree *Tree, relation *json.RelationCommon, i int) (*Union, error) {
	union := Union{
		Comments: relation.Comments,
	}

	switch relation.Type {
	case "wedding":
		union.Kind = UnionWedding
	case "civil":
		union.Kind = UnionCivil
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
