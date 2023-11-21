Gair aenerate migration

```bash
migrate create -seq -ext=.sql -dir=./migrations name
```

To run migrations on the backend

```bash
migrate -path . -database "postgres://goexpense_user:goexpense@localhost/expense_db?sslmode=disable" up
```
