package dictionary

type Dictionary map[string]string

func Create() *Dictionary {
	d := make(Dictionary)
	return &d
}

func (d Dictionary) Add(word, translate string) error {
	d[word] = translate
	return nil
}

func (d Dictionary) Delete(word string) error {
	delete(d, word)
	return nil
}

func (d Dictionary) ReturnInfo() map[string]string {
	newMap := make(map[string]string)
	for word, result := range newMap {
		newMap[word] = result
	}
	return newMap
}
