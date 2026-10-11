package dictionary

type Dictionary map[string]string

func Create() *map[string]string {
	d := make(map[string]string)
	return &d
}

func (d Dictionary) Add(word, translate string) error {
	d[word] = translate
	return nil
}

//доделтаь вывод списка, подумав про то, как сделать так, чтобы вывод словаря был независимым
