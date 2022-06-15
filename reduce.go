package collection

func (m enumerable) Reduce(fn interface{}, memo interface{}) IEnumerable {
	return m.Aggregate(fn, memo)
}
