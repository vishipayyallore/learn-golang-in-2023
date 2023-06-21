# Sample Bank Application in GoLang

Some Description

## Set the Environment Variables

> 1. "root:Sample@123$@tcp(localhost:3306)/banking"

```powershell
$Env:SERVER_ADDRESS="localhost"
$Env:SERVER_PORT=8080
$Env:DB_USER="root"
$Env:DB_PASSWD="Sample@123$"
$Env:DB_ADDR="localhost"
$Env:DB_PORT="3306"
$Env:DB_NAME="banking"
```

## Subtitle

Some Description

```powershell
PS C:\LordKrishna\GitHub\learn-golang-in-2023\bankappdemo\src\services> go mod init bankappdemo/services
go: creating new go.mod: module bankappdemo/services
go: to add module requirements and sums:
        go mod tidy
PS C:\LordKrishna\GitHub\learn-golang-in-2023\bankappdemo\src\services> go mod tidy
PS C:\LordKrishna\GitHub\learn-golang-in-2023\bankappdemo\src\services> cd ../..
PS C:\LordKrishna\GitHub\learn-golang-in-2023\bankappdemo> go work use .\src\services\
PS C:\LordKrishna\GitHub\learn-golang-in-2023\bankappdemo> go get github.com/jmoiron/sqlx

go work use .\src\dtos\
go work sync

```



> 1. <https://github.com/gorilla/mux>
> 1. <https://github.com/go-sql-driver/mysql>
> 1. [go get github.com/jmoiron/sqlx](go get github.com/jmoiron/sqlx)
