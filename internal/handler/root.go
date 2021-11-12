package handler

import (
	"github.com/spf13/cobra"
)

func (h *handlers) Root(cmd *cobra.Command, args []string) {
	cobra.CheckErr(cmd.Usage())
}
