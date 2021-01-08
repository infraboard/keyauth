package types

// Is todo
func (t *UserType) Is(tps ...UserType) bool {
	for _, tp := range tps {
		if *t == tp {
			return true
		}
	}

	return false
}
