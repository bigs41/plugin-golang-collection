package collection

type enumerable struct {
	Enumerator func() IEnumerator
}

func (m enumerable) GetEnumerator() IEnumerator {
	return m.Enumerator()
}
