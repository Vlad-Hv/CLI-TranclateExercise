package cli

import (
	"bufio"
	"exerciser/internal/handler/ui"
	"exerciser/internal/service"
)

func dictionaryMenu(s service.Service, in *bufio.Scanner) error {
	for {
		option, err := ui.DictionaryMenu(in)

		if err != nil {
			return err
		}

		if option == 4 {
			ui.LeftMenu()
			break
		}

		switch option {
		case 1:
			w, t, err := ui.NewPair(in)

			if err != nil {
				return err
			}

			if err = s.NewPair(w, t); err != nil {
				return err
			}

			ui.AddSuccess()

		case 2:
			w, err := ui.DeleteWord(in)

			if err != nil {
				return err
			}

			if err = s.DeletePair(w); err != nil {
				return err
			}

			ui.DeleteSuccess()

		case 3:
			list := s.Storage.ReturnInfo()
			ui.PrintWordList(list)

		default:
			ui.UnexistOption()
		}
	}
	return nil
}

func Training(s service.Service, amount int, in *bufio.Scanner) error {
	result := make(map[string]string)
	if err := service.ValidateAmount(amount); err != nil {
		return err
	}

	if err := s.Storage.ValidateWordsAmount(amount); err != nil {
		return err
	}

	var counter int
	var correct int
	words := s.Storage.ReturnInfo()

	for word, translate := range words {
		userResponce, err := ui.GetTranslate(word, in)
		if err != nil {
			return err
		}
		if service.CheckAnswer(translate, userResponce) {
			result[word] = "correct"
			correct++
			ui.Correct()
		} else {
			result[word] = "uncorrect"
			ui.Uncorrect()
		}
		counter++
		if counter >= amount {
			break
		}
	}

	err := s.SaveResult(result)
	if err != nil {
		return err
	}

	ui.PrintResult(amount, correct)
	return nil
}
