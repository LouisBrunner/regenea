package genea

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/LouisBrunner/regenea/genea/json"
)

func Parse(in io.Reader) (*Tree, Version, error) {
	content, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, 0, err
	}

	var hdr json.Header
	err = json.Parse(content, &hdr)
	if err != nil {
		return nil, 0, err
	}

	var tree *Tree
	var version Version

	switch version = Version(*hdr.Version); version {
	case VersionV1:
		var jsonRepr json.V1
		err = json.Parse(content, &jsonRepr)
		if err != nil {
			return nil, 0, err
		}
		tree, err = ImportV1(&jsonRepr)

	case VersionV2:
		var jsonRepr json.V2
		err = json.Parse(content, &jsonRepr)
		if err != nil {
			return nil, 0, err
		}
		tree, err = ImportV2(&jsonRepr)

	default:
		return nil, 0, fmt.Errorf("unknown version: `%d`", version)
	}

	return tree, version, err
}

func initTree(comments string, people int) *Tree {
	return &Tree{
		Comments: comments,
		People:   make([]*Person, people),
		ByID:     make(map[string]*Person),
	}
}

func createPerson(jsonPerson *json.PersonCommon) *Person {
	person := Person{
		ID:       string(jsonPerson.ID),
		Sex:      SexOther,
		Comments: jsonPerson.Comments,
	}
	switch jsonPerson.Sex {
	case "M":
		person.Sex = SexMale
	case "F":
		person.Sex = SexFemale
	}
	return &person
}

func addPerson(tree *Tree, i int, person *Person) error {
	if tree.ByID[person.ID] != nil {
		return fmt.Errorf("duplicate person ID: `%s`", person.ID)
	}
	tree.People[i] = person
	tree.ByID[person.ID] = person
	return nil
}

func getPerson(tree *Tree, personID *json.PersonID) (*Person, error) {
	if personID == nil || string(*personID) == "" {
		return nil, nil
	}
	id := string(*personID)
	person := tree.ByID[id]
	if person == nil {
		return nil, fmt.Errorf("unknown person ID: `%s`", id)
	}
	return person, nil
}

func listHasPerson(haystack []*Person, needle *Person) bool {
	for _, person := range haystack {
		if person == needle {
			return true
		}
	}
	return false
}
