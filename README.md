# cloudfile CLI

Cloudfiles CLI is a go program take advantage of `github.com/ncw/swift` to perform
operation against cloudfiles (racksapce).

## Env variable

This program looks in the shell environ to find the `UserName` and the `ApiKey`
required to connect to cloudfiles.

```
export UserName="MyUserName"
export ApiKey="MySecretApiKey"
```

## Setting up

## Bulk upload a tar.gz into a container

```
Cloudfiles -timeout 60s - container <my-container> -file <the-file.tar.gz>
```
