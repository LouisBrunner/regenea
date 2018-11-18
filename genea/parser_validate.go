package genea

import (
	"fmt"
)

func compareEvents(event1 *EventCore, event2 *EventCore) bool {
	if !EventIsValid(event1) || !EventIsValid(event2) {
		return true
	}
	return event1.Date.Unix() <= event2.Date.Unix()
}

func (tree *Tree) validate() error {
	// TODO: check for mismatch between death of parent and birth of child
	for _, person := range tree.People {
		if !compareEvents(&person.Birth, person.Death) {
			return fmt.Errorf("person `%s`: birth cannot be after death", person.ID)
		}
	}
	return tree.IterateOverUnionErr(func(union *Union) error {
		if !compareEvents(&union.Begin, union.End) {
			return fmt.Errorf("relation between `%s` and '%s': start cannot be after end", union.Person1, union.Person2)
		}
		return nil
	})
}
