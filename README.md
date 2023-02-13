# FastTubeAnalytics

![Build status](https://github.com/daniil49926/FastTubeAnalytics/actions/workflows/healthchecker.yml/badge.svg?branch=main)

API service for analytics of the [FastTube](https://github.com/daniil49926/FastTube).
The service is designed to collect and issue analytics data for the service that is being serviced. These are assembled using different handles.

## Technology stack

- Go 1.19
- Postgresql
- Docker

## Before run

Before starting work, you need to create a toml file containing the data necessary to start the server.

*File location: ./configs/prod/api.toml*

***Configuration file variable***
```toml
bind_addr = ""
log_level = ""
postgres_dsn = ""
```

## Build
```
go build cmd/api/main.go
```

For the service, assembly is provided in a lightweight docker container, using multistage build.

## Build and Run in Docker container
```
docker build --tag fast-tube-analytics:1 .
docker-compose up
```


## Description of methods

<details>
<summary> List of available endpoints </summary>

1. A simple endpoint for health check:

**Request**
```
GET /healthchecker
```
**Response**
```json
{
  "result": "OK"
}
```

2. Request to send analytics:

**Request**
```
GET /send-analitics
```
The request accepts a raw string as input, of the form: "REQUEST-method-url-process_time;RESPONSE-status"  

**Response**
```json
{
  "result": "OK"
}
```

3. Request to get all the analytics.

**Request**
```
GET /get-analytics
```
**Response**
```json
{
  "result": [
    ["1","GET","1.1.1.1:1111","0.01111","200"],
    [...], 
  ]
}
```

The request returns all the collected analytics data

</details>

