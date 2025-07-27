package ansible

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/galaxy/collection/install"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

type AnsibleCollection struct {
	Name string
}

func (c AnsibleCollection) RunPlaybook(playbookName string) error {
	// Executes one of the playbook contained in the collection.
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "localhost,",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(fmt.Sprintf("%s.%s", c.Name, playbookName)),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		execute.WithEnvVars(map[string]string{
			"ANSIBLE_FORCE_COLOR": "True",
			//"ANSIBLE_CONFIG":      "/etc/ansible/ansible.cfg",
		}),
	)

	return exec.Execute(context.TODO())
}

func (c AnsibleCollection) InstallFromZip(file []byte) error {
	// Installs the collection, as provided by the `file` (zip) file.
	// This zip file can be an embed.

	// Step 1: Write embedded ZIP to disk as a temporary file
	tmpZip, err := os.CreateTemp("/tmp", "{{ cookiecutter.cli_binary_name }}_collection_")
	if err != nil {
		panic(err)
	}

	err2 := os.WriteFile(tmpZip.Name(), file, 0644)
	if err2 != nil {
		panic(err2)
	}
	defer os.Remove(tmpZip.Name()) // Ensure removal

	// Step 2: Install colelction, pointing to zip
	collectionInstallCmd := galaxycollectioninstall.NewAnsibleGalaxyCollectionInstallCmd(
		galaxycollectioninstall.WithCollectionNames(tmpZip.Name()),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(collectionInstallCmd),
	)

	return exec.Execute(context.TODO())
}
