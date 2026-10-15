package service

import (
	"exerciser/internal/storage"
	"fmt"
)

type Service struct {
	Storage storage.WordStorage
	Results storage.ResultStorage
}

func (s Service) NewPair(word, translate string) error {
	if err := s.Storage.ValidateExist(word); err != nil {
		return fmt.Errorf("cannot make new word and translate pair: %w", err)
	}

	return s.Storage.Add(word, translate)
}

func (s Service) DeletePair(word string) error {
	if err := s.Storage.ValidateUnexist(word); err != nil {
		return err
	}
	if err := s.Storage.Delete(word); err != nil {
		return err
	}
	return nil
}

func NewWordStorage(storage storage.WordStorage, result storage.ResultStorage) *Service {
	return &Service{Storage: storage, Results: result}
}

//в обработчик добавить меню, пересобрать функции сервсиа в цикле меню
