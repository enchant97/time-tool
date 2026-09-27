package core

func DefaultIfUnset[T comparable](v, d, u T) T {
	if v == u {
		return d
	}
	return v
}
