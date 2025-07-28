# {{ cookiecutter.ansible_collections_slug }}

## Description
This project contains Ansible collections to perform workstation configuration tasks, in the specific context of the company.

## Project Structure
- `{{ cookiecutter.ansible_collections_namespace }}`: the company's namespace.
- `pyproject.toml`: To generate a venv containing development tools.
- `mise.toml`: Versioning of development tooling (your workstation comes equipped with it ;) )

## Installation
To install the collections, use the following command:
```bash
ansible-galaxy collection install {{ cookiecutter.ansible_collections_namespace }}.{{ cookiecutter.ansible_collections_name }}
```

## Contributing
To contribute to this project, please refer to the [CONTRIBUTING](CONTRIBUTING) file.
