package genea

import (
	"encoding/json"
)

func (tree *Tree) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"People": tree.People,
	})
}

func litePerson(person *Person) interface{} {
	if person == nil {
		return nil
	}
	return map[string]interface{}{
		"ID":    person.ID,
		"Name":  person.Name(),
		"Sex":   person.Sex,
		"Birth": person.Birth,
		"Death": person.Death,
	}
}

func liteUnion(union *Union) interface{} {
	if union == nil {
		return nil
	}
	return map[string]interface{}{
		"Kind":    union.Kind,
		"Person1": litePerson(union.Person1),
		"Person2": litePerson(union.Person2),
		"Begin":   union.Begin,
		"End":     union.End,
	}
}

func (person *Person) marshalJSON(pretty bool) ([]byte, error) {
	data := map[string]interface{}{
		"ID":     person.ID,
		"Name":   person.Name(),
		"Sex":    person.Sex,
		"Birth":  person.Birth,
		"Death":  person.Death,
		"Father": litePerson(person.Father),
		"Mother": litePerson(person.Mother),
		"Family": liteUnion(person.Family),
	}
	if pretty {
		return json.MarshalIndent(data, "", "  ")
	}
	return json.Marshal(data)
}

func (person *Person) MarshalJSON() ([]byte, error) {
	return person.marshalJSON(false)
}

func (sex Sex) MarshalJSON() ([]byte, error) {
	var str string
	switch sex {
	case SexMale:
		str = "Male"
	case SexFemale:
		str = "Female"
	case SexOther:
		str = "Other"
	default:
		str = "Unknown"
	}
	return []byte("\"" + str + "\""), nil
}
