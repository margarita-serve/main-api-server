package cmd

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/cmd/db"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/cmd/server"
)

func init() {
	RootCmd.AddCommand(db.DBCmd)
	RootCmd.AddCommand(server.ServerCmd)
}
