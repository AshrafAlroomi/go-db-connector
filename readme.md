
# Go lang test project

A db connector and rest-api





## Run Locally

Install go lang
```bash
   sudo snap install go --classic
```

[Install postgresql](https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart)



Clone the project

```bash
  git clone https://github.com/AshrafAlroomi/go-db-connector
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
  go get .
```

Start the server

```bash
   source .env && go run main.go
```


## Running Tests

To run tests, run the following command

Add a new user
```bash
  curl --location --request POST 'http://localhost:8007/users' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "name":"ashraf",
    "email":"ashraf@code-care.pro"
}'
```
response
```bash
success 200
{
    "Value": {
        "ID": 4,
        "CreatedAt": "2023-02-20T16:35:34.951881065+03:00",
        "UpdatedAt": "2023-02-20T16:35:34.951881065+03:00",
        "DeletedAt": null,
        "Name": "ashraf2",
        "Email": "ashrasf2@code-care.pro"
    },
    "Error": null,
    "RowsAffected": 1
}
```

Get all the users 

```bash
curl --location --request GET 'http://localhost:8007/users' \
--header 'Content-Type: application/json'
```
response
```bash
success 200
[
    {
        "ID": 1,
        "CreatedAt": "2023-02-18T15:44:39.23706Z",
        "UpdatedAt": "2023-02-18T15:44:39.23706Z",
        "DeletedAt": null,
        "Name": "ashraf",
        "Email": "ashraf@code-care.pro"
    },
    {
        "ID": 2,
        "CreatedAt": "2023-02-18T19:58:30.295301Z",
        "UpdatedAt": "2023-02-18T19:58:30.295301Z",
        "DeletedAt": null,
        "Name": "ashraf2",
        "Email": "ashraf2@code-care.pro"
    },
    {
        "ID": 4,
        "CreatedAt": "2023-02-20T13:35:34.951881Z",
        "UpdatedAt": "2023-02-20T13:35:34.951881Z",
        "DeletedAt": null,
        "Name": "ashraf2",
        "Email": "ashrasf2@code-care.pro"
    }
]
```
delete user
```bash
curl --location --request DELETE 'http://localhost:8007/users/1' \
--header 'Content-Type: application/json'
```
response
```bash
success 200
```

update user
```bash
curl --location --request PUT 'http://localhost:8007/users/2' \
--header 'Content-Type: application/json'
--data-raw '{
    "name":"ashraf2",
    "email":"ashrafNEW@code-care.pro"
}
```
response
```bash
success 200

{
    "Value": {
        "ID": 2,
        "CreatedAt": "2023-02-20T16:35:34.951881065+03:00",
        "UpdatedAt": "2023-02-20T16:35:34.951881065+03:00",
        "DeletedAt": null,
        "Name": "ashraf2",
        "Email": "ashrafNEW@code-care.pro"
    },
    "Error": null,
    "RowsAffected": 1
}

```


Migrate to db
```bash
curl --location --request GET 'http://localhost:8007/migrate' \
--header 'Content-Type: application/json'
```
response
```bash
success 200```