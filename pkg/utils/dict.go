package utils

type Dict[K comparable, V any] map[K]V

func NewDict[K comparable, V any](values ...map[K]V) Dict[K, V] {
	if len(values) == 0 {
		return make(Dict[K, V])
	}
	return values[0]
}

func (d Dict[K, V]) Set(key K, value V) {
	d[key] = value
}

func (d Dict[K, V]) Get(key K) V {
	return d[key]
}

func (d Dict[K, V]) Has(key K) bool {
	_, ok := d[key]
	return ok
}

func (d Dict[K, V]) Delete(key K) {
	delete(d, key)
}

func (d Dict[K, V]) Keys() []K {
	keys := make([]K, 0, len(d))
	for key := range d {
		keys = append(keys, key)
	}
	return keys
}

func (d Dict[K, V]) Values() []V {
	values := make([]V, 0, len(d))
	for _, value := range d {
		values = append(values, value)
	}
	return values
}

func (d Dict[K, V]) Len() int {
	return len(d)
}

func (d Dict[K, V]) Clear() {
	for key := range d {
		delete(d, key)
	}
}

func (d Dict[K, V]) Copy() Dict[K, V] {
	copy := make(Dict[K, V])
	for key, value := range d {
		copy[key] = value
	}
	return copy
}

func (d Dict[K, V]) Merge(other Dict[K, V]) {
	for key, value := range other {
		d[key] = value
	}
}

func (d Dict[K, V]) Iterate(fn func(key K, value V)) {
	for key, value := range d {
		fn(key, value)
	}
}

func (d Dict[K, V]) Map(fn func(key K, value V) V) Dict[K, V] {
	result := make(Dict[K, V])
	for key, value := range d {
		result[key] = fn(key, value)
	}
	return result
}

func (d Dict[K, V]) Filter(fn func(key K, value V) bool) Dict[K, V] {
	result := make(Dict[K, V])
	for key, value := range d {
		if fn(key, value) {
			result[key] = value
		}
	}
	return result
}

func (d Dict[K, V]) Reduce(fn func(key K, value V, accum V) V, initial V) V {
	accum := initial
	for key, value := range d {
		accum = fn(key, value, accum)
	}
	return accum
}
