package embedded

import _ "embed"

//go:embed {{ cookiecutter.ansible_collections_namespace }}-{{ cookiecutter.ansible_collections_name }}-1.0.0.tar.gz
var CollectionZip []byte
