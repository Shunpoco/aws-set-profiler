package prompt

import (
	"github.com/manifoldco/promptui"
)

func GetProfile(items []string) (result string, err error) {
	prompt := promptui.Select{
		Label: "Select AWS Profile",
		Items: items,
	}

	_, result, err = prompt.Run()

	return
}
