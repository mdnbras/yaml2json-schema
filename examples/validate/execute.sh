./jsonschema.exe generate --yaml examples/validate/input.yaml --csv examples/validate/metadata.csv --output examples/validate/output.schema.json --verbose --dump-normalized-yaml-file examples/validate/normalized.yaml --remove-anchor-bases pipeline-base

./jsonschema.exe validate --schema examples/validate/output.schema.json --input examples/validate/normalized.yaml