package cli

import (
	"exerciser/internal/handler/ui"
	"exerciser/internal/service"
)

func dictionaryMenu(s service.Service) error {
	for {
		option, err := ui.DictionaryMenu()

		if err != nil {
			return err
		}

		if option == 4 {
			ui.LeftMenu()
			break
		}

		switch option {
		case 1:
			w, t, err := ui.NewPair()

			if err != nil {
				return err
			}

			if err = s.NewPair(w, t); err != nil {
				return err
			}

			ui.AddSuccess()

		case 2:
			w, err := ui.DeleteWord()

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
