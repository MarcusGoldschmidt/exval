# Exval

API able to evaluate arbitrary logical expressions.

## Api

User Basic auth to authenticate.

Default user is **admin** and password is **admin**

Swagger is available at **/swagger**

## Expression

An expression is a string that can be evaluated to a boolean value.

The current implementation

```
S                   -> EXPRESSION CONTINUE_EXPRESSION
EXPRESSION          -> ID CONTINUE_EXPRESSION | ( EXPRESSION )
CONTINUE_EXPRESSION -> OPERATOR EXPRESSION CONTINUE_EXPRESSION | &
OPERATOR            -> AND | OR
```

An alternative to this implementation is to use Reverse polish notation.

### Building the application

To run the tasks below you need, **docker** and **go** instaled

```shell
# create api binary
make api
# setup postgres
docker-compose up -d postgres
# Execute aplication
# The default port is 11139
./api
```

### Tests

```shell
go test ./...
```

## Contribute

1. Fork the repository
2. Create a pull request
3. Wait for review
4. Merge
5. Celebrate
6. Repeat XD