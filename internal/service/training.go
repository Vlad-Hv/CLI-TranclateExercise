package service

import (
	"exerciser/internal/handler/ui"
	"strings"
)

func (s Service) Training(amount int) error {
	if err := s.Storage.ValidateWordsAmount(amount); err != nil {
		return err
	}
	result := make(map[string]string)
	var counter int
	words := s.Storage.ReturnInfo()

	for word, translate := range words {
		userResponce := ui.GetTranslate(word)
		t, r := modifyWords(translate, userResponce)

		if strings.EqualFold(t, r) {
			result[word] = "correct"
			ui.Correct()
		} else {
			result[word] = "mistake"
			ui.Uncorrect()
		}
		counter++

		if counter >= amount {
			break
		}
	}

	err := s.Results.Add(result)
	if err != nil {
		return err
	}

	return nil
}

func modifyWords(t, r string) (string, string) {
	translate := strings.TrimSpace(t)
	responce := strings.TrimSpace(r)

	return translate, responce
}
