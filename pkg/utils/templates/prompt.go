package templates

import "github.com/manifoldco/promptui"

func NewPromptSelect(label string, items []string) (string, error) {
	 prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()
	return result, err
}
