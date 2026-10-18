package storage

type WordStorage interface {
	Add(word, translate string) error
	Delete(word string) error
	Validator
	Returner
}

type Validator interface {
	ValidateExist(word string) error
	ValidateUnexist(word string) error
	ValidateWordsAmount(amount int) error
}

type Returner interface {
	ReturnInfo() map[string]string
}

type ResultStorage interface {
	Add(result map[string]string) error
	ReturnResults() []map[string]string
}
