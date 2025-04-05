package hashmap

// FilterKeys returns a new map after filtering the keys that satisfy the filter function
func FilterKeys[K comparable, V any](m map[K]V, filter func(K) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if filter(k) {
			result[k] = v
		}
	}
	return result
}

// SelectKeys returns a new map after selecting the keys that are in the keys slice
func SelectKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	return FilterKeys(m, func(k K) bool {
		for _, key := range keys {
			if k == key {
				return true
			}
		}
		return false
	})
}
