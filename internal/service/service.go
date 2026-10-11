package service

import (
	"exerciser/internal/handler/ui"
	"exerciser/internal/storage"
	"fmt"
)

type Service struct {
	Storage storage.Storage
}

func (s Service) NewPair() error {
	w, t, err := ui.NewPair()

	if err != nil {
		return err
	}

	if err = s.Storage.ValidateExist(w); err != nil {
		return fmt.Errorf("cannot make new word and translate pair: %w", err)
	}

	err = s.Storage.Add(w, t)
	if err != nil {
		return err
	} else {
		ui.AddSuccess()
	}

	return nil
}
