# ImageScraper

## Build
### In Windows

```powershell
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip # If you do not this have utility earlier
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o .\Handler .\main.go # from directory where main.go or handler code is present
~\Go\Bin\build-lambda-zip.exe -o .\Handler.zip .\Handler
```

```bash
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip # If you do not this have utility earlier
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o .\Handler .\main.go # considering your are in the same directory where your main.go or handler code is present
~\Go\Bin\build-lambda-zip.exe -o .\Handler.zip .\Handler
```

### Linux & Mac

```bash
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

## Using Lambda
In lambda
```
{
    "link": "link to image to download"
    "key": "name the image to be download"
}
```
run the function