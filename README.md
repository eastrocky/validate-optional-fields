# validate-optional-fields
A demonstration on validating user input with optional fields

## Running the Server
```zsh
git clone https://github.com/eastrocky/validate-optional-fields.git

go build

./validate-optional-fields
```


## Invalid favoriteNumber
### Request
```zsh
curl -X POST \
-d '{"favoriteNumber": 12}' \
localhost:8080/payload
```
### Response
```json
{
    "code":400,
    "message":"Key: 'Payload.FavoriteNumber' Error:Field validation for 'FavoriteNumber' failed on the 'max' tag",
    "status":"Bad Request"
}
```

## Valid favoriteNumber
### Request
```zsh
curl -X POST \
-d '{"favoriteNumber": 2}' \
localhost:8080/payload
```

### Response
```json
{
    "code":200,
    "message":"Your favoriteNumber is 2",
    "status":"OK"
}
```

## Omitting favoriteNumber
### Request
```zsh
curl -X POST \
-d '{}' \
localhost:8080/payload
```

### Response
```json
{
    "code":200,
    "message":"You did not provide favoriteNumber",
    "status":"OK"
}
```