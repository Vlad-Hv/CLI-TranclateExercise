package storage

type Storage interface {
	Add(word, translate string) error
	//Delete
	Validator
}

type Validator interface {
	ValidateExist(word string) error
}
