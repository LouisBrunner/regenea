package genea

import (
	"fmt"
)

// TODO: add way to get subset of a tree (all parents, siblings (child included) and children)

func addPersonAndParents(tree *Tree, person *Person) *Person {
	_, found := tree.ByID[person.ID]
	if found {
		return nil
	}
	newPerson := *person

	// TODO: redo that
	newPerson.Children = []*Person{}
	newPerson.Family = nil
	newPerson.Partners = []*Union{}

	newPerson.Siblings = []*Person{}

	tree.People = append(tree.People, &newPerson)
	tree.ByID[person.ID] = &newPerson
	if person.Father != nil {
		newPerson.Father = addPersonAndParents(tree, person.Father)
	}
	if person.Mother != nil {
		newPerson.Mother = addPersonAndParents(tree, person.Mother)
	}
	return &newPerson
}

func (tree *Tree) Subtree(id string) (*Tree, error) {
	person, found := tree.ByID[id]
	if !found {
		return nil, fmt.Errorf("could not find person `%s`", id)
	}

	newTree := initTree(tree.Comments, 0)
	addPersonAndParents(newTree, person)
	return newTree, nil
}
