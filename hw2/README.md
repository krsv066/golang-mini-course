API
=

1. Создать аккаунт
```
curl 0.0.0.0:8000/account -H "Content-Type: application/json" -X POST --data '{"name": "ACCOUNT_NAME"}'    
```

2. Получить аккаунт
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X GET
```

3. Изменить баланс аккаунта
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X PATCH --data '{"balance": INT}'
```

4. Удалить аккаунт
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X DELETE
```

5. Изменить имя аккаунта
```
curl 0.0.0.0:8000/account/{ACCOUNT_NAME} -H "Content-Type: application/json" -X PUT --data '{"name": "NEW_ACCOUNT_NAME"}'
```

CLI
=

```
go build -o client
```

1. Создать аккаунт
```
./client create --name "ACCOUNT_NAME" 
```

2. Получить аккаунт
```
./client get --name "ACCOUNT_NAME"
```

3. Изменить баланс аккаунта
```
./client update --name "ACCOUNT_NAME" --balance INT
```

4. Удалить аккаунт
```
./client delete --name "ACCOUNT_NAME"
```

5. Изменить имя аккаунта
```
./client rename --name "ACCOUNT_NAME" --new_name "NEW_ACCOUNT_NAME"
```
