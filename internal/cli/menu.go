package cli

import (
	"exerciser/internal/handler/ui"
	"exerciser/internal/service"
)

func Menu(s service.Service) (bool, error) {
	option, err := ui.MenuOption()

	if err != nil {
		return false, err
	}

	if option == 4 {
		return true, nil
	}
	switch option {
	case 1:
		err := dictionaryMenu(s)
		if err != nil {
			return false, err
		}
	case 2:

	case 3:

	default:
		ui.UnexistOption()
	}
	return false, nil
}
