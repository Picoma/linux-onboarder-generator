# Ansible Collection - {{ cookiecutter.ansible_collections_namespace }}.{{ cookiecutter.ansible_collections_name }}

Documentation for the collection.

## Description
This project contains Ansible collections to perform workstation configuration tasks, in the specific context of the company.

## Installation
To install this collection, use the following command:
```bash
ansible-galaxy collection install {{ cookiecutter.ansible_collections_namespace }}.{{ cookiecutter.ansible_collections_name }}
```

## Usage
Here is an example of using this collection in a playbook:
```yaml
---
- name: Example playbook using the collection
  hosts: all
  collections:
    - {{ cookiecutter.ansible_collections_namespace }}.{{ cookiecutter.ansible_collections_name }}
  tasks:
    - name: Example task
      {{ cookiecutter.ansible_collections_namespace }}.{{ cookiecutter.ansible_collections_name }}.module_name:
        param1: value1
        param2: value2
```

You can also directly call the playbooks:
```bash
ansible-playbook -i 'localhost,' -c local {{ cookiecutter.ansible_collections_namespace }}.{{ cookiecutter.ansible_collections_name }}.init
```

## Documentation
For more details on the available modules and plugins, refer to the official documentation.

## Contributing
To contribute to this collection, please refer to the CONTRIBUTING file.