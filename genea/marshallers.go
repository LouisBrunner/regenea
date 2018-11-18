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
		"Names": person.Name(),
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

func (person *Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"ID":     person.ID,
		"Names":  person.Name(),
		"Sex":    person.Sex,
		"Birth":  person.Birth,
		"Death":  person.Death,
		"Father": litePerson(person.Father),
		"Mother": litePerson(person.Mother),
		"Family": liteUnion(person.Family),
	})
}
