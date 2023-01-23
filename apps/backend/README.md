# backend

Projet for backend app written in go.

## Prerequisites

[Traefik](https://gitlab.com/florenttorregrosa-docker/apps/docker-traefik) for working locally with it.

### Swagger

#### Comment your code

Decorate controllers following this [doc](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format)

#### Generate swagger docs

Using https://github.com/swaggo/gin-swagger

`go install github.com/swaggo/swag/cmd/swag`

then run

`swag init`
