# Schema Validator Generator

CLI em Go para gerar JSON Schema (Draft-07) a partir de um YAML de estrutura e um CSV de metadados, e validar um YAML contra o schema gerado.

## Visao geral
- **generate**: combina YAML + CSV e gera um `schema.json` valido no Draft-07.
- **validate**: valida um `input.yaml` contra um `schema.json`.
- Suporta normalizacao de YAML com anchors, diff e dump do YAML normalizado.

## Requisitos
- Go instalado (ver `BUILD.md` para exemplos de build)

## Build
```shell
go build -o jsonschema ./cmd/jsonschema
go build -o jsonschema.exe ./cmd/jsonschema
```

## Uso rapido
### Gerar schema
```shell
./jsonschema.exe generate --yaml examples/simple/input.yaml --csv examples/simple/metadata.csv --output examples/simple/output.schema.json --verbose
```

### Validar YAML contra schema
```shell
./jsonschema.exe validate --schema examples/validate/output.schema.json --input examples/validate/normalized.yaml
```

## Comandos e flags
### Comando raiz
- `--verbose`, `-v`: habilita logs detalhados.

### generate
Gera JSON Schema Draft-07 combinando YAML e CSV.

Flags:
- `--yaml`, `-y`: caminho do arquivo YAML de entrada (obrigatorio)
- `--csv`, `-c`: caminho do arquivo CSV com metadados (obrigatorio)
- `--output`, `-o`: caminho do JSON Schema de saida (default: `schema.json`)
- `--schema-version`: versao do JSON Schema (default: `https://json-schema.org/draft-07/schema#`)
- `--yaml-diff`: exibe diff entre YAML original e YAML normalizado
- `--dump-normalized-yaml-file`: salva o YAML normalizado no arquivo informado
- `--remove-anchor-bases`: lista de chaves (separadas por virgula) removidas apos resolver anchors

### validate
Valida um YAML contra um JSON Schema Draft-07.

Flags:
- `--schema`: caminho do arquivo `schema.json` (obrigatorio)
- `--input`: caminho do arquivo `input.yaml` (obrigatorio)

## Formato do CSV de metadados
Cabecalho esperado:
```
path,description,required,type
```

Regras principais (exemplos em `examples/simple/metadata.csv`):
- **Caminho por pontos** para propriedades: `user.name`
- **Array** com `[]`: `pipelines.ingestion[]`
- **Map/dicionario** com `*`: `pipelines.ingestion.*.enabled`

Exemplo:
```csv
path,description,required,type
user.name,Nome do usuario,true,string
pipelines.ingestion.*,Pipeline,true,object
pipelines.ingestion.*.enabled,Habilita pipeline,true,boolean
```

## Exemplos prontos
- `examples/simple`: geracao basica a partir de YAML + CSV.
- `examples/anchors`: exemplo com anchors, diff e YAML normalizado.
- `examples/validate`: geracao + validacao do YAML normalizado.

Comandos usados nos exemplos:
```shell
./jsonschema.exe generate --yaml examples/anchors/input.yaml --csv examples/anchors/metadata.csv --output examples/anchors/output.schema.json --verbose --yaml-diff --dump-normalized-yaml-file examples/anchors/normalized.yaml --remove-anchor-bases pipeline-base

./jsonschema.exe generate --yaml examples/validate/input.yaml --csv examples/validate/metadata.csv --output examples/validate/output.schema.json --verbose --dump-normalized-yaml-file examples/validate/normalized.yaml --remove-anchor-bases pipeline-base
./jsonschema.exe validate --schema examples/validate/output.schema.json --input examples/validate/normalized.yaml
```

## Links uteis
- Converter YAML para JSON: https://onlineyamltools.com/convert-yaml-to-json
- Validar JSON/YAML contra schema: https://www.jsonschemavalidator.net/
