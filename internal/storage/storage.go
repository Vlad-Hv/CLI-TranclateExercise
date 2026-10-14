package storage

type Storage interface {
	Add(word, translate string) error
	Delete(word string) error
	Validator
	Returner
}

type Validator interface {
	ValidateExist(word string) error
	ValidateUnexist(word string) error
}

type Returner interface {
	ReturnInfo() map[string]string
}
