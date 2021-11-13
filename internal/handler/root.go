package handler

import (
	"github.com/spf13/cobra"
)

type rootHandler struct{}

func (h *rootHandler) Run(cmd *cobra.Command, args []string) {
	cobra.CheckErr(cmd.Usage())
}
