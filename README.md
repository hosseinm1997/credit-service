# Credit service part of "Arvan Cloud" challenge

This service is responsible for handling credit codes.

## Quick start
To quickly jump into the main logic go to following links:

- [Spend credit code API logic](https://github.com/hosseinm1997/credit-service/blob/ab1eda279aa9e2a4d02b4d752e09de0e0f3da42f/http/endpoints/SpendCodeEndpoint.go#L17)
- [Inquiry credit code API logic](https://github.com/hosseinm1997/credit-service/blob/ab1eda279aa9e2a4d02b4d752e09de0e0f3da42f/http/endpoints/SpendCodeEndpoint.go#L71)

## Overview

### Approach

The most important challenge was overcoming race conditions. My approach to solving this problem is creating a **database function** (PostgreSQL DB engine) that implemented **`Optimistic Concurrency Control`** (Database Optmistic Lock).

See [utilize_credit_code()](https://github.com/hosseinm1997/credit-service/blob/main/db/routines/functions/postgres/utilize_credit_code.sql)

> **Note: For better performance, I use pgbouner(connection pooling) over PostgreSQL.**

### Architecture
I Use the dual write strategy for communication between these two microservices. It's better to use distributed transaction management patterns, especially the Saga pattern via the Orchestrator model. 

### Framework
This service was made based on a simple framework made by myself (in a limited time). I'm not interested in `reinvent the wheel` myself!! My idea behind this is to dig into the Go language deeper. It has following features:

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

### Framework TODOs:
- [ ] Use gRPC for communication
- [ ] Use Saga pattern instead of dual write strategy
- [ ] Segregate reads and writes via CQRS pattern 
- [ ] Use gorm migration
- [ ] Use unit testing with high code coverage
- [ ] Pass context into internal services
- [ ] Use swagger for API documentation

### Credit Service TODOs:
- [ ] Make buffered queue before requesting to DB function

<br/>
<br/>
<br/>

## Credit Service Docs

This service mainly focuses on how to spend credit codes. It has two APIs. One for **inquiring** about credit code availability & usability, and another for **spending** codes.


### - Inquiry API:
- **signature**: `/credit/code/{code}/inquiry` 
- **inputs**:

    `{code}` (string): credit code text receive from `wallet service`.

- **description**: This endpoint looks for the `{code}` credit code in `codes` table. Then check the result whether the credit is available and current used count is less then maximum usable count.


### - Spending API:
- **signature**: `/credit/code/{code}/{referenceId}` 
- **inputs**:

    `{code}` (string): credit code text receive from `wallet service`.

    `{referenceId}` (integer): An integer value using for keeping track of successful operation between two services. Will be Generated and sent by `wallet service`

- **description**: This endpoint at first check for code availabity and usability. If it was ok, Then it calls `utilize_credit_code()` database function Afterward it process thendatabase result then respond to the `wallet service`.


