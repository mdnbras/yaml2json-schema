./jsonschema.exe generate --yaml examples/validate/json/input.yaml --csv examples/validate/json/metadata.csv --output examples/validate/json/output.schema.json --verbose --dump-normalized-yaml-file examples/validate/json/normalized.yaml --remove-anchor-bases pipeline-base

./jsonschema.exe validate --schema examples/validate/json/output.schema.json --input examples/validate/json/input.json