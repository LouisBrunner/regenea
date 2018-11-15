package genea

func (tree *Tree) IterateOverUnion(do func(union *Union)) {
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
				do(person.Family)
				person.Family.flags.serialized = true
			}
		}
		for _, partner := range person.Partners {
			if !partner.flags.serialized {
				do(partner)
				partner.flags.serialized = true
			}
		}
	}
}
