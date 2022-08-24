# Credit service part of "Arvan Cloud" challenge

This service is responsible for handling credit codes.

## Quick start
To quickly jump into the main logic go to following links:

- [Spend credit code API logic](https://github.com/hosseinm1997/credit-service/blob/ab1eda279aa9e2a4d02b4d752e09de0e0f3da42f/http/endpoints/SpendCodeEndpoint.go#L17)
- [Inquiry credit code API logic](https://github.com/hosseinm1997/credit-service/blob/ab1eda279aa9e2a4d02b4d752e09de0e0f3da42f/http/endpoints/SpendCodeEndpoint.go#L71)

## Overview

### Approach

The most important challenge was overcoming race conditions. My approach to solving this problem is using a **database function** (Postgres DB engine) that implemented **`Optimistic Concurrency Control`** (Database Optmistic Lock).

### Framework
This service was made based on a simple framework made by myself (in a limited time) that has these features:

- IoC implemented using service container, created by new `generic` feature of go 1.18. [see ServiceContainer.go](https://github.com/hosseinm1997/credit-service/blob/main/infrastructures/ServiceContainer.go)
- Routing system using middlewares and contextes. [see RoutingSystem.go](https://github.com/hosseinm1997/credit-service/blob/main/infrastructures/RoutingSystem.go)
- Easy exception handling with `Respond()` helper function. [see ResponseFormatter.go](https://github.com/hosseinm1997/credit-service/blob/main/http/middlewares/ResponseFormatter.go), [see an example](https://github.com/hosseinm1997/credit-service/blob/ab1eda279aa9e2a4d02b4d752e09de0e0f3da42f/http/endpoints/SpendCodeEndpoint.go#L71)
- Handling env variables
- Service and repository pattern considered

### Packages
Direct packages used:

- Viper for managing env variables
- Go-chi for routing
- Gorm for using as an ORM and query builder


This service mainly focuses on how to spend credit codes. It has two APIs. One for **inquiring** about credit code availability & usability, and another for **spending** codes.
### - Inquiry API:
- **signature**: `/credit/code/{code}/{referenceId}` 
- **inputs**:

    `{code}` (string): credit code text receive from `wallet service`.

    `{referenceId}` (integer): An integer value using for keeping track of successful operation between two services. Will be Generated and sent by `wallet service`

- **description**: This endpoint looks for the `{code}` credit code in `codes` table. Then check the result whether the credit is available and current used count is less then maximum usable count.


