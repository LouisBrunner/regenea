package genea

func (tree *Tree) IterateOverUnionErr(do func(union *Union) error) error {
	for _, person := range tree.People {
		if person.Family != nil {
			person.Family.flags.serialized = false
		}
		for _, partner := range person.Partners {
			partner.flags.serialized = false
		}
	}

	for _, person := range tree.People {
		if person.Family != nil {
			if !person.Family.flags.serialized {
				err := do(person.Family)
				if err != nil {
					return err
				}
				person.Family.flags.serialized = true
			}
		}
		for _, partner := range person.Partners {
			if !partner.flags.serialized {
				err := do(partner)
				if err != nil {
					return err
				}
				partner.flags.serialized = true
			}
		}
	}

	return nil
}

func (tree *Tree) IterateOverUnion(do func(union *Union)) {
	tree.IterateOverUnionErr(func(union *Union) error {
		do(union)
		return nil
	})
}
