### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```shell
migrate -database "postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable" -path db/migrations up
```