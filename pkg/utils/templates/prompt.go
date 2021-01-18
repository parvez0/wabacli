package templates

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func NewPromptSelect(label string, items []string) (string, error) {
	 prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()
	return result, err
}

func NewPromptPassword() (string, error) {
	validate := func(input string) error {
		if len(input) < 1 {
			return fmt.Errorf("password must not be empty")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}
	return prompt.Run()
}
