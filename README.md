#

[pre-commit](https://pre-commit.com/) is used to manage pre-commit hooks.

- Install pre-commit: `pip install pre-commit`
- Install hooks: `pre-commit install`

#### download
```shell
go mod download
```
#### run docker
```shell
docker-compose -f deploy/docker-compose-local.yaml up -d
```
#### run main
```shell
go run .
```
