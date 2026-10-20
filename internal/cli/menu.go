package cli

import (
	"bufio"
	"exerciser/internal/handler/ui"
	"exerciser/internal/service"
)

func Menu(s service.Service, in *bufio.Scanner) (bool, error) {
	const programStop bool = true
	const programContinue bool = false
	option, err := ui.MenuOption(in)

	if err != nil {
		return programContinue, err
	}

	if option == 4 {
		return programStop, nil
	}
	switch option {
	case 1:
		err := dictionaryMenu(s, in)
		if err != nil {
			return programContinue, err
		}
	case 2:
		wordsAmount, err := ui.GetAmoutWords(in)
		if err != nil {
			return programContinue, err
		}

		err = Training(s, wordsAmount, in)
		if err != nil {
			return programContinue, err
		}

	case 3:
		results := s.Results.ReturnResults()
		ui.PrintResults(results)

	default:
		ui.UnexistOption()
	}
	return programContinue, nil
}
