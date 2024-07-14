Run Docker and Postgres
=

```
docker run -p 5432:5432 --name some-postgres -e POSTGRES_PASSWORD=password -d postgres
psql -h 0.0.0.0 -p 5432 -U postgres
CREATE TABLE accounts(name varchar(256) PRIMARY KEY, balance int not null default 0);
```

API
=

1. Create account
```
curl 0.0.0.0:8000/account -H "Content-Type: application/json" -X POST --data '{"name": "ACCOUNT_NAME"}'    
```

2. Get account
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X GET
```

3. Rename account
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X PUT --data '{"name": "NEW_ACCOUNT_NAME"}'
```

4. Update account balance
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X PATCH --data '{"balance": INT}'
```

5. Delete account
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X DELETE
```

CLI
=

```
go build -o client
```

1. Create account
```
./client create --name "ACCOUNT_NAME" 
```

2. Get account
```
./client get --name "ACCOUNT_NAME"
```

3. Rename account
```
./client rename --name "ACCOUNT_NAME" --new_name "NEW_ACCOUNT_NAME"
```

4. Update account balance
```
./client update --name "ACCOUNT_NAME" --balance INT
```

5. Delete account
```
./client delete --name "ACCOUNT_NAME"
```
