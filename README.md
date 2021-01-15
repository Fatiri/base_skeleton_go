# The XAMPLE API
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/makes-people-smile.svg)](https://forthebadge.com)

## Genesis :blue_heart: Golang
 
 ## Requirements
 
  - Golang 
  - Go Module
  - other needed modules
  
 ## Installing
  
   - Use Go Modules, please read https://blog.golang.org/using-go-modules
   - Install Dependecies
      ```console
      $ go get
      ```
   - Copy file env.example to be .env
   - Fill your local configuration
   - Run Project (go run .)
 
## Migration files

Each migration has an up and down migration.

```bash
1481574547_create_accounts_table.up.sql
1481574547_create_accounts_table.down.sql
```

#### Why two separate files (up and down) for a migration?
  It makes all of our lives easier. No new markup/syntax to learn for users
  and existing database utility tools continue to work as expected.

#### How many migrations can migrate handle?
  Whatever the maximum positive signed integer value is for your platform.
  For 32bit it would be 2,147,483,647 migrations. Migrate only keeps references to
  the currently run and pre-fetched migrations in memory.

[Best practices: How to write migrations.](MIGRATIONS.md)
 
 Create File Migration
 
 ```console
$ go run migration/create.go file_name 
 ```

#### How to run db migration using migrate?

Through Docker run 
```console 
docker run -v $DIRECTORY_SCHEMA:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME?multiStatements=true" up 
```

Through Binary 
```console
1. Download binary from here https://github.com/golang-migrate/migrate/releases/download/v4.2.4/migrate.linux-amd64.tar.gz 
2. Extract binary and copy binary to /opt/migrate.linux-amd64
3. Run migration --> /opt/migrate.linux-amd64 -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME?multiStatements=true" -source file://$DIRECTORY_SCHEMA/ up 
```

## Coding Style Guide

See this [guide](https://github.com/uber-go/guide/blob/master/style.md) for better code.

## Database 
Please read this source for better performance
 - [Use the index luke](https://use-the-index-luke.com/) 
 - [GORM Performance](https://gorm.io/docs/performance.html)

## Logger
We use [Zap](https://github.com/uber-go/zap), performance benchmarking 

Log a message and 10 fields:

| Package             |    Time     | Time % to zap | Objects Allocated |
| :------------------ | :---------: | :-----------: | :---------------: |
| :zap: zap           |  862 ns/op  |      +0%      |    5 allocs/op    |
| :zap: zap (sugared) | 1250 ns/op  |     +45%      |   11 allocs/op    |
| zerolog             | 4021 ns/op  |     +366%     |   76 allocs/op    |
| go-kit              | 4542 ns/op  |     +427%     |   105 allocs/op   |
| apex/log            | 26785 ns/op |    +3007%     |   115 allocs/op   |
| logrus              | 29501 ns/op |    +3322%     |   125 allocs/op   |
| log15               | 29906 ns/op |    +3369%     |   122 allocs/op   |

Log a message with a logger that already has 10 fields of context:

| Package             |    Time     | Time % to zap | Objects Allocated |
| :------------------ | :---------: | :-----------: | :---------------: |
| :zap: zap           |  126 ns/op  |      +0%      |    0 allocs/op    |
| :zap: zap (sugared) |  187 ns/op  |     +48%      |    2 allocs/op    |
| zerolog             |  88 ns/op   |     -30%      |    0 allocs/op    |
| go-kit              | 5087 ns/op  |    +3937%     |   103 allocs/op   |
| log15               | 18548 ns/op |    +14621%    |   73 allocs/op    |
| apex/log            | 26012 ns/op |    +20544%    |   104 allocs/op   |
| logrus              | 27236 ns/op |    +21516%    |   113 allocs/op   |

Log a static string, without any context or `printf`-style templating:

| Package             |    Time    | Time % to zap | Objects Allocated |
| :------------------ | :--------: | :-----------: | :---------------: |
| :zap: zap           | 118 ns/op  |      +0%      |    0 allocs/op    |
| :zap: zap (sugared) | 191 ns/op  |     +62%      |    2 allocs/op    |
| zerolog             |  93 ns/op  |     -21%      |    0 allocs/op    |
| go-kit              | 280 ns/op  |     +137%     |   11 allocs/op    |
| standard library    | 499 ns/op  |     +323%     |    2 allocs/op    |
| apex/log            | 1990 ns/op |    +1586%     |   10 allocs/op    |
| logrus              | 3129 ns/op |    +2552%     |   24 allocs/op    |
| log15               | 3887 ns/op |    +3194%     |   23 allocs/op    |

#### How to use logger? 
The logger format will be written like this
``` console
{"level":"ERROR","time":"2020-10-31T05:32:29.580+0700","caller":github.com/base_skeleton_go/src/repository/account.go:66","message":"AccountRepository-GetByEmail Error record not found"}
```

- Error Log With Format
``` console
    logger.Ef(` SetFirstLogin %s `, err) 
```

- Error Log Without Format
``` console
    logger.E(err) 
```

- Info Log With Format
``` console
    logger.If(` SetFirstLogin %s `, &user) 
```

- Info Log Without Format
``` console
    logger.I(&use) 
```

You can find another logger method at shared/logger package
 

 ## Data Validation
For validation we use [ozzo-validation](https://github.com/go-ozzo/ozzo-validation).
is a Go package that provides configurable and extensible data validation capabilities.
It has the following features:

- use normal programming constructs rather than error-prone struct tags to specify how data should be validated.
- can validate data of different types, e.g., structs, strings, byte slices, slices, maps, arrays.
- can validate custom data types as long as they implement the `Validatable` interface.
- can validate data types that implement the `sql.Valuer` interface (e.g. `sql.NullString`).
- customizable and well-formatted validation errors.
- provide a rich set of validation rules right out of box.
- extremely easy to create and use custom validation rules.

## Example

```go
data := "base"
if err := validation.Validate(data, validation.Required); err != nil {
		fmt.Println(err)
	}
}

if err := validation.Validate(data, validation.Required, validation.Length(5, 100)); err != nil {
		fmt.Println(err)
	}
}
```

 ## Environment Variables
 
 ```
DEVELOPMENT=
PORT=

DB_HOST=
DB_USERNAME=
DB_PASSWORD=
DB_NAME=
DB_PORT=

TOKEN_NAME="X-base-AccessToken"
TOKEN_TYPE=Bearer
TOKEN_TEMPORARY_LIFETIME=300
ACCESS_TOKEN_LIFETIME=120

