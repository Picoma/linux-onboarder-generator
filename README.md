# Linux-onboarder-generator

## Description
This Cookiecutter template is designed to simplify and accelerate the onboarding of Linux users in your company. With this project, you can generate a robust and customized workstation configuration utility, ensuring a smooth and efficient user experience.

## Features
- 🚀 Simplified Deployment
    - **CLI in Go**: A single generated binary, easy to deploy on workstations, offering optimal performance.
    - **Cobra Framework**: Using Cobra for rapid CLI development and comfortable autocompletion script generation for users.

- 🔧 Improved Developer Experience
    - **Mise Integration**: Bootstrapped with [`mise`](https://mise.jdx.dev/) to enrich the developer experience through its tool, task, and environment management. The CLI itself is managed by mise!

- 🌐 Extensibility and Control
    - **Internal Playbooks**: Based on internal company playbooks, allowing you to stay 100% in control of the infrastructure.

## Project Structure
Two repositories are generated:
- **`{{ cookiecutter.ansible_collections_slug }}`**: Contains internal Ansible collections for workstation configuration.
- **`{{ cookiecutter.cli_tool_slug }}`**: A CLI in Go with Cobra, ready to use.

## Usage
To get started, initialize a project with Cookiecutter:
```bash
cookiecutter https://github.com/Picoma/linux-onboarder-generator
```

Then, you will need to:
1. Build the collection (a mise task is provided for this) in `.zip` format.
2. Move this `.zip` artifact into the `embedded` folder of the CLI project, and ensure the `embeds.go` file embeds the correct file.
3. Build the CLI; again, a mise task is ready to use.

Refer to the documentation of the generated projects for detailed instructions.
