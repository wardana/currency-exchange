# Currency Exchange #
Currency Exchange is a simple application for storing daily exchange data and display foreign exchange rate for currencies on a daily basis.

## Prerequisites: ##
- Golang
    -  Download: https://golang.org/dl/
- Docker
    - Download: https://www.docker.com/get-started
- MySQL Docker Image
    - Download https://hub.docker.com/_/mysql/
- Docker Compose
    - Download https://docs.docker.com/compose/install/#install-compose   

notes:  make sure mysql in docker expose your mysql port and also for application we need to expose speficy port

## Usecases: ##
- User wants to input daily exchange rate data
- User has a list of exchange rates to be tracked
- User wants to add an exchange rate to the list
- User wants to remove an exchange rate from the list
- User wants to see the exchange rate trend from the most recent 7 data points

# Exposed Endpoint #
`GET /api/v1/currency`
```
## Purposes
Get all available currency pair

## Success Response
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 11 Aug 2018 09:44:32 GMT
Content-Length: n
{
  "status": 200,
  "message": "Success",
  "data": [
    {
      "id": 2,
      "base_currency": "IDR",
      "counter_currency": "USD",
      "created_at": "2018-08-09T12:31:08Z",
      "updated_at": "2018-08-09T12:31:08Z"
    }
  ]
}
```

`POST /api/v1/currency/create`
```
## Purposes
Added new currency pair

## Http Headers
Content-Type: application/json

## Sampe Payload
{
  "base_currency": "IDR",
  "counter_currency": "SGD"
}

## Success Response
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 11 Aug 2018 09:44:09 GMT
Content-Length: n

{
  "status": 201,
  "message": "Data has been created",
  "data": {
    "id": 7,
    "base_currency": "IDR",
    "counter_currency": "GBR",
    "created_at": "2018-08-11T09:44:09.494814473Z",
    "updated_at": "2018-08-11T09:44:09.494814473Z"
  }
}
```
`PUT /api/v1/currency/:id`
```
## Purposes
Edit currency pair data using id as parameter

## Http Headers
Content-Type: application/json

## Sample Payload 
{
  "base_currency": "IDR",
  "counter_currency": "MYR"
}

## Success Response
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 11 Aug 2018 09:46:52 GMT
Content-Length: n

{
  "status": 200,
  "message": "Success",
  "data": {
    "id": 7,
    "base_currency": "IDR",
    "counter_currency": "MYR",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  }
}
```
`DELETE /api/v1/currency/:id`
```
## Purposes
Remove listed currency pair using id as parameter

## Success Response
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 11 Aug 2018 09:50:36 GMT
Content-Length: n

{
  "status": 200,
  "message": "Success",
  "data": {}
}1
```

`POST /api/v1/rate`
```
## Purposes
Added new daily exchange rate data for specify currency pair

## Http Headers
Content-Type: application/json

## Sample Payload
{
  "base_currency": "IDR",
  "counter_currency": "JPY",
  "exchange_rate": 0.5, // float
  "exchange_date": "2018-08-11" //format 'YYYY-MM-DD'
}

## Success Response
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 11 Aug 2018 09:55:53 GMT
Content-Length: n

{
  "status": 201,
  "message": "Data has been created",
  "data": {
    "exchange_rate": 0.5,
    "exchange_date": "2018-08-11T00:00:00Z",
    "base_currency": "IDR",
    "counter_currency": "JPY",
    "weekly_average": 0
  }
}
```
`GET /api/v1/rate/exchange`
```
## Purpose
Get list of exchange rates

## Url Parameter
date= YYY-MM-DDD
example: ?date=2018-08-11

or leave it blank for set current date as parameter

## Http Headers
Content-Type: application/json

## Success Response
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 11 Aug 2018 10:03:13 GMT
Content-Length: n

{
  "status": 200,
  "message": "Success",
  "data": [
    {
      "exchange_rate": 0.4,
      "exchange_date": "2018-08-11T09:44:09.494814473Z",
      "base_currency": "IDR",
      "counter_currency": "USD",
      "weekly_average": 0.4
    }
  ]
}

```

`GET /api/v1/rate/trend`
```
## Purpose
Get exchange rate trend from the most recent 7 data points

## Url Parameter
base= currency_id
counter= currency_id
example: ?base=IDR&counter=USD

## Success Response
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Sat, 11 Aug 2018 10:06:49 GMT
Content-Length: 574

{
  "status": 200,
  "message": "Success",
  "data": {
    "base_currency": "IDR",
    "counter_currency": "USD",
    "average": 0.31428571428571433,
    "variance": 0.4,
    "history": [
      {
        "exchange_rate": 0.1,
        "exchange_date": "2018-08-10T07:00:00Z"
      },
      {
        "exchange_rate": 0.2,
        "exchange_date": "2018-08-09T07:00:00Z"
      }
    ]
  }
}
```

###

### Run Your Applicaiton ###
```
$ docker-compose up

------------------------------------
Enable log mode         - false
Environment             - development
------------------------------------
[INFO] Connected to MySQL. Config => db, LogMode => false

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.5
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
 ⇨ http server started on [::]:8080
```
