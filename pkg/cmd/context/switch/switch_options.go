package _switch

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"regexp"
)

const (
	ListAllAccounts string = "list_all"
	PartialMatch string = "partial_match"
	NoChange string = "no_change"
)

type SwitchOptions struct {
	Config *config.Configuration
	Type string
	Query string
}

func (s *SwitchOptions) Validate(c []string)  {
	if len(c) > 1 {
		log.Error("")
	}
	if len(c) == 0 {
		s.Type = ListAllAccounts
		return
	}
	s.Query = c[0]
	if s.Config.CurrentContext == c[0] || s.Config.CurrentCluster.Name == c[0] {
		log.Debug(fmt.Sprintf("context '%s' is already active", c[0]))
		s.Type = NoChange
		return
	}
	s.Type = PartialMatch
	return
}

func (s *SwitchOptions) Run()  {
	switch s.Type {
	case NoChange:
		return
	case ListAllAccounts:
		s.listAll()
	case PartialMatch:
		s.partialMatch()
	}
}

// listAll will list all the accounts as interactive
// shell, where you can select the account
func (s *SwitchOptions) listAll() {
	clus := make([]string, len(s.Config.Clusters))
	for i, c := range s.Config.Clusters {
		clus[i] = c.Context
	}
	log.Debug("found the following context in config: ", clus)
	result, err := s.prompt(clus)
	if err != nil {
		log.Error("prompt failed with error: ", err)
	}
	log.Debug("activating selected context '" + result + "'")
	err = config.UpdateContext(result)
	if err != nil {
		log.Error("failed to update context: ", err)
	}
}

// partialMatch will take the query and regex match
// with the existing accounts, if the account not
// found it will exit with error code 1
func (s *SwitchOptions) partialMatch()  {
	// creating a regexp for the query provided
	reg, err := regexp.Compile(fmt.Sprintf(`%s`, s.Query))
	if err != nil {
		log.Error("failed to create regex exp: ", err)
	}
	var accs []string

	// getting all the matched accounts from the cluster
	for _, v := range s.Config.Clusters {
		if matched := reg.MatchString(v.Context); matched {
			accs = append(accs, v.Context)
		}
	}
	var result string
	if len(accs) > 1 {
		result, err = s.prompt(accs)
		if err != nil {
			log.Error("prompt failed with error: ", err)
		}
	}

	if len(accs) == 0 {
		log.Error(fmt.Sprintf("account with name '%s', not found", s.Query))
	}

	if len(accs) == 1 {
		if accs[0] == s.Query {
			result = accs[0]
		} else {
			result, err = s.prompt(accs)
			if err != nil {
				log.Error("prompt failed with error: ", err)
			}
		}
	}

	log.Debug("activating selected context '" + result + "'")
	err = config.UpdateContext(result)
	if err != nil {
		log.Error("failed to update context: ", err)
	}
}

// prompt will show multiple selectable options to the user
func (s *SwitchOptions) prompt(items []string) (string, error) {
	log.Debug("multiple accounts detected, ", items)
	prompt := promptui.Select{
		Label: "Select Account",
		Items: items,
	}
	_, result, err := prompt.Run()
	return result, err
}


