package handler

import "github.com/spf13/cobra"

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) Run(cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
