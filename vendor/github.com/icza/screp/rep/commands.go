// This file contains the types describing the players' commands.

package rep

import "github.com/icza/screp/rep/repcmd"

// Commands contains the players' commands.
type Commands struct {
	// Cmds is the commands of the players
	Cmds []repcmd.Cmd

	// ParseErrCmds is list of commands that failed to parse.
	// A parse error command may imply additional skipped (not recorded) commands
	// at the same frame.
	ParseErrCmds []*repcmd.ParseErrCmd
}
