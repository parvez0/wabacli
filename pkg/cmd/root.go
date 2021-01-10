package cmd

import "github.com/spf13/cobra"

func NewDefaultCmd() *cobra.Command {
	return wabacliCmd()
}

func wabacliCmd() *cobra.Command {
	return &cobra.Command{}
}

