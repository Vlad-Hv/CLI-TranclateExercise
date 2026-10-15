package result

type ResultStorage []map[string]string

func NewStorage() ResultStorage {
	var results []map[string]string
	return results
}

func (rs *ResultStorage) Add(result map[string]string) error {
	*rs = append(*rs, result)
	return nil
}

func (rs ResultStorage) ReturnResults() []map[string]string {
	return rs
}
