package change

import (
	"fmt"
	"github.com/parvez0/wabacli/config"
	"github.com/parvez0/wabacli/log"
	"github.com/parvez0/wabacli/pkg/errutil/handler"
	"github.com/parvez0/wabacli/pkg/utils/templates"
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

func NewSwitchOptions(c *config.Configuration) *SwitchOptions {
	return &SwitchOptions{
		Config: c,
	}
}

func (s *SwitchOptions) Validate(c []string) error {
	if len(c) > 1 {
		return fmt.Errorf("select a single account")
	}
	if len(c) == 0 {
		s.Type = ListAllAccounts
		return nil
	}
	s.Query = c[0]
	if s.Config.CurrentContext == c[0] || s.Config.CurrentCluster.Name == c[0] {
		log.Debug(fmt.Sprintf("context '%s' is already active", c[0]))
		s.Type = NoChange
		return nil
	}
	s.Type = PartialMatch
	return nil
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
	result, err := templates.NewPromptSelect("Select Account", clus)
	if err != nil {
		handler.FatalError(fmt.Errorf("prompt failed with error: %v", err))
	}
	log.Debug("activating selected context '" + result + "'")
	err = s.Config.UpdateContext(result)
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to update context: %v", err))
	}
}

// partialMatch will take the query and regex match
// with the existing accounts, if the account not
// found it will exit with error code 1
func (s *SwitchOptions) partialMatch()  {
	l, accs := s.Config.MatchAcc(s.Query)
	var result string
	var err error
	if l > 1 {
		result, err = templates.NewPromptSelect("Select Account", accs)
		if err != nil {
			handler.FatalError(fmt.Errorf("prompt failed with error: %v", err))
		}
	}

	if l == 0 {
		handler.FatalError(fmt.Errorf("account with name '%s', not found", s.Query))
	}

	if l == 1 {
		if accs[0] == s.Query {
			result = accs[0]
		} else {
			result, err = templates.NewPromptSelect("Select Account", accs)
			if err != nil {
				handler.FatalError(fmt.Errorf("prompt failed with error: %v", err))
			}
		}
	}

	log.Debug("activating selected context '" + result + "'")
	err = s.Config.UpdateContext(result)
	if err != nil {
		handler.FatalError(fmt.Errorf("failed to update context: %v", err))
	}
}


