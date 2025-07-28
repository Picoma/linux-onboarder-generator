#!/bin/bash

set -euo pipefail

SCRIPT_DIR=$(dirname $0)
PYTHON_VERSION=${PYTHON_VERSION:-latest}

echo "Bootstrapping {{ cookiecutter.cli_binary_name }} installation."

echo "Installing mise :"
curl --silent --show-error https://mise.run/bash | MISE_QUIET=1 sh
eval "$(~/.local/bin/mise --quiet activate)"
echo

echo "Bootstrapping userland Python and ansible through mise"
mise --quiet use -g python@${PYTHON_VERSION}
eval "$(~/.local/bin/mise activate)" # Activate now to add python to path.

python -m pip --quiet install --user pipx
mise --quiet use -g ansible@11

mise --quiet settings experimental=true
# Post-generation tip : you can change this http backend to github or gitlab to automatically pull new releases. See https://mise.jdx.dev/dev-tools/backends/ 
mise --quiet use -g http:{{ cookiecutter.cli_binary_name }}[url=http://localhost:8000/{{ cookiecutter.cli_binary_name }}]@1.0.0

eval "$(~/.local/bin/mise activate)" # Re-activate add {{ cookiecutter.cli_binary_name }} to path
{{ cookiecutter.cli_binary_name }} init

echo "\nInstallation terminée ! La fenêtre va maintenant se fermer. Bon code !"
