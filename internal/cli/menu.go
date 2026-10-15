package cli

import (
	"exerciser/internal/handler/ui"
	"exerciser/internal/service"
)

func Menu(s service.Service) (bool, error) {
	const programStop bool = true
	const programContinue bool = false
	option, err := ui.MenuOption()

	if err != nil {
		return programContinue, err
	}

	if option == 4 {
		return programStop, nil
	}
	switch option {
	case 1:
		err := dictionaryMenu(s)
		if err != nil {
			return programContinue, err
		}
	case 2:
		wordsAmount, err := ui.GetAmoutWords()
		if err != nil {
			return programContinue, err
		}
		err = s.Training(wordsAmount)
		if err != nil {
			return programContinue, err
		}
		//success message

	case 3:
		results := s.Results.ReturnResults()
		ui.PrintResults(results)

	default:
		ui.UnexistOption()
	}
	return programContinue, nil
}
