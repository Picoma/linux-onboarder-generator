/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"{{ cookiecutter.cli_package_name }}/embedded"
	"{{ cookiecutter.cli_package_name }}/runners"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the workstation configuration.",
	Long: `Initializes the workstation configuration :

Installs the embedded corporate Ansible collection, and runs its 'init' playbook.

Requires Ansible installed.
`,
	Run: func(cmd *cobra.Command, args []string) {
		c := &ansible.AnsibleCollection{
			Name: "{{ cookiecutter.ansible_collections_namespace }}.{{ cookiecutter.ansible_collections_name }}",
		}

		c.InstallFromZip(embedded.CollectionZip)
		c.RunPlaybook("init")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
