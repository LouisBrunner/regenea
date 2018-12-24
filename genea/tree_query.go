package genea

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

func (tree *Tree) Query(ctxt *cli.Context) (interface{}, error) {
	requestedID := ctxt.String("id")
	if requestedID != "" {
		person, found := tree.ByID[requestedID]
		if found {
			bytes, err := person.marshalJSON(true)
			if err != nil {
				return nil, err
			}
			return string(bytes), nil
		}
		return nil, fmt.Errorf("Person not found for ID `%s`", requestedID)
	}
	return nil, fmt.Errorf("no criteria provided")
}
