/*
  Copyright (c) 2018, Sylabs, Inc. All rights reserved.

  This software is licensed under a 3-clause BSD license.  Please
  consult LICENSE file distributed with the sources of this project regarding
  your rights to use or distribute this software.
*/

package cli

import (
	"github.com/singularityware/singularity/docs"
	"github.com/spf13/cobra"
)

var capabilityUse string = `capability <subcommand>`

var capabilityShort string = `Manage Linux capabilities`

var capabilityLong string = `
Capabilities allow you to have fine grained control over the permissions that 
your containers need to run. For instance, if you need to `

var capabilityExample string = `
All group commands have their own help output:

$ singularity help capability add
$ singularity capability list --help
`

func init() {

	manHelp := func(c *cobra.Command, args []string) {
		docs.DispManPg("singularity-capability")
	}

	SingularityCmd.AddCommand(CapabilityCmd)
	CapabilityCmd.SetHelpFunc(manHelp)
	CapabilityCmd.AddCommand(CapabilityAddCmd)
	CapabilityCmd.AddCommand(CapabilityDropCmd)
	CapabilityCmd.AddCommand(CapabilityListCmd)
}

var CapabilityCmd = &cobra.Command{
	Run: nil,
	DisableFlagsInUseLine: true,

	Use:     capabilityUse,
	Short:   capabilityShort,
	Long:    capabilityLong,
	Example: capabilityExample,
}
