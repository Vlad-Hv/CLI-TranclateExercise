package result

type ResultStorage []map[string]string

func NewStorage() ResultStorage {
	var results []map[string]string
	return results
}

func (rs *ResultStorage) Add(result map[string]string) error {
	copied := copyResult(result)
	*rs = append(*rs, copied)
	return nil
}

func (rs ResultStorage) ReturnResults() []map[string]string {
	n := make([]map[string]string, len(rs))
	for i := range rs {
		result := copyResult(rs[i])
		n[i] = result
	}
	return n
}

func copyResult(r map[string]string) map[string]string {
	results := make(map[string]string)
	for word, result := range r {
		results[word] = result
	}
	return results
}
