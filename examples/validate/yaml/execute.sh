./jsonschema.exe generate --yaml examples/validate/yaml/input.yaml --csv examples/validate/yaml/metadata.csv --output examples/validate/yaml/output.schema.json --verbose --dump-normalized-yaml-file examples/validate/yaml/normalized.yaml --remove-anchor-bases pipeline-base

./jsonschema.exe validate --schema examples/validate/yaml/output.schema.json --input examples/validate/yaml/normalized.yaml