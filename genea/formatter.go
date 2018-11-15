package genea

import (
	"fmt"
	"io"

	"github.com/LouisBrunner/regenea/genea/json"
	"github.com/json-iterator/go"
)

func (tree *Tree) Format(out io.Writer, pretty bool, version Version) error {
	var data interface{}
	switch version {
	case VersionV1:
		data = serializeV1(tree)
	case VersionV2:
		data = serializeV2(tree)
	default:
		return fmt.Errorf("unsupported version `%v`", version)
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var content []byte
	var err error
	if pretty {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}
	n, err := out.Write(content)
	if err != nil {
		return err
	}
	if n != len(content) {
		return fmt.Errorf("incomplete write")
	}
	return nil
}

func serializeRelationCommon(union *Union) json.RelationCommon {
	relation := json.RelationCommon{}
	switch union.Kind {
	case UnionWedding:
		relation.Type = "wedding"
	case UnionCivil:
		relation.Type = "civil"
	}
	if union.Person1 != nil {
		id := json.PersonID(union.Person1.ID)
		relation.Person1 = &id
	}
	if union.Person2 != nil {
		id := json.PersonID(union.Person2.ID)
		relation.Person2 = &id
	}
	relation.Comments = union.Comments
	return relation
}

func serializePersonCommon(person *Person) json.PersonCommon {
	sex := "?"
	switch person.Sex {
	case SexMale:
		sex = "M"
	case SexFemale:
		sex = "F"
	}
	return json.PersonCommon{
		ID:       json.PersonID(person.ID),
		Sex:      sex,
		Comments: person.Comments,
	}
}
