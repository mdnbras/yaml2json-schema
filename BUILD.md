- BUILD
```shell
go build -o jsonschema ./cmd/jsonschema
go build -o jsonschema.exe ./cmd/jsonschema
```

- RUN
```shell
./jsonschema.exe generate --yaml examples/input.yaml --csv examples/metadata.csv --output examples/output.schema.json --verbose
```

- EXAMPLE CSV METADATA
```csv
path,description,required,type
user.name,Nome do usuário,true,string
user.age,Idade do usuário,false,number
user.emails,Emails,false,array
pipelines.ingestion,Lista de pipelines,true,array
pipelines.ingestion[],Pipeline de ingestão,true,object
pipelines.ingestion[].enabled,Habilita pipeline,true,boolean
pipelines.ingestion[].description,Descrição,false,string
pipelines.ingestion[].source,Configuração da origem,true,object
pipelines.ingestion[].source.connection-ref,Conexão,true,string
pipelines.ingestion[].source.query,Query SQL,true,string
pipelines.ingestion[].staging,Configuração staging,true,object
pipelines.ingestion[].staging.collection-name,Coleção,true,string
pipelines.ingestion,Map de pipelines,true,object
pipelines.ingestion.*,Pipeline,true,object
pipelines.ingestion.*.enabled,Habilita pipeline,true,boolean
pipelines.ingestion.*.description,Descrição,false,string
pipelines.ingestion.*.source,Configuração da origem,true,object
pipelines.ingestion.*.source.connection-ref,Conexão,true,string
pipelines.ingestion.*.source.query,Query SQL,true,string
pipelines.ingestion.*.staging,Configuração staging,true,object
pipelines.ingestion.*.staging.collection-name,Coleção,true,string
```