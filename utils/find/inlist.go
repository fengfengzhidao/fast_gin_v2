package find

func InList[T comparable](list []T, key T) (ok bool) {
	for _, t := range list {
		if t == key {
			return true
		}
	}
	return
}
