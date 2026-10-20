package service

import (
	"strings"
)

func CheckAnswer(translate, userResponce string) bool {
	t, ur := modifyWords(translate, userResponce)

	if strings.EqualFold(t, ur) {
		return true
	} else {
		return false
	}
}

func modifyWords(t, r string) (string, string) {
	translate := strings.TrimSpace(t)
	responce := strings.TrimSpace(r)

	return translate, responce
}

func (s *Service) SaveResult(result map[string]string) error {
	err := s.Results.Add(result)
	if err != nil {
		return err
	}

	return nil
}
