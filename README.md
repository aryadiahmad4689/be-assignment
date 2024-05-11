# Wallet

## Overvie2
  This application certainly has many shortcomings, therefore I only handle a small part of the features that can be done here. I just want to give an overview of how the test provided works. I hope it's easy to understand.

## Features
- **User Authentication**: Supports sign up and login functionalities.
- **Send Balance**: Users can make transfers between accounts
- **Withdraw**: Users can make withdrawals from their accounts


## Tech Stack
- **Language**: [golang] i choose becauser is faster, mature and big community.
- **Database**: [postgres] - i choose because is recomended
- **Managemen** : [makefile] i choose because is easily to run and migrate,this also complies with https://12factor.net/admin-processes
- **Envirotment**: [.env] i choose is easily manage envirotment,this also complies with https://12factor.net/config
- **Sofware Architecture** : [layered architecture] I chose it because it is easy to develop and is modular and meets the principles of clean architecture https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html


## Getting Started

### Prerequisites
- Golang 1.18 up
- VsCode
- Postman

### Detail And Structure
- `cmd`: Main of service
- `migrate` for migrate database
- `src`: Content and logic of the application
  - `app`: Main configuration settings of the application
    - `db`: To create configuration and connection to the database
  - `middleware`: To create middleware and manage people who log in to the application
  - `modules`: Business logic and core application

### Dependencies
- github.com/dgrijalva/jwt-go v3.2.0+incompatible
- github.com/go-chi/chi v1.5.5
- github.com/go-chi/cors v1.2.1
- github.com/jabardigitalservice/golog v0.0.8
- github.com/joho/godotenv v1.5.1
- github.com/mattn/go-sqlite3 v1.14.19
- github.com/spf13/viper v1.18.2
- github.com/stretchr/testify v1.8.4
- golang.org/x/crypto v0.16.0
- gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
- gopkg.in/go-playground/assert.v1 v1.2.1

### Installation And Runnin Not Docker
1. Clone the repository:
   ```bash
   git clone https://github.com/aryadiahmad4689/be-assignment.git 
  
2. Enter the main folder 

3. Copy the `.env.example ` to `.env`

4. Field full envirotment

5. Run `make migrate` this is to migrate database `note: must to have postgres`

6. Run `go mod tidy` to get dependency

7. Run `make run` to running apps

8. Congrats app is up


### Installation And Runnin Docker Compose
1. Clone the repository:
   ```bash
   git clone https://github.com/aryadiahmad4689/be-assignment.git 
  
2. Enter the main folder 

3. Copy the `.env.example ` to `.env`

4. Field full envirotment and customize with docker compose

5. Run `docker compose up`

6. Congrats app is up

### Url Api And Method
 - `/v1/account/sign-up` method `post`
 - `/v1/account/sign-in` method `post`
 - `/v1/account/get-all-user-account` method `get`
 - `/v1/payment/send` Method `post`
 - `/v1/payment/widthraw` Method `post`


### Postman
You can import file json `collection_test.json` to postman.
- i only handle positif test
- before you test in postman you must to create envirotmen and add variable `AUTH_USER`


### Testing
For unit test i just handle two layer (repository and usecase). I decided not to handle endpoint and handler because of time constraints. Please understand
 - Unit test run `make test`
 - Coverage uni test run `make test-coverage`

### Linting
 - VsCode go extension


### Advice
I understand that my system definitely has a lot of gaps and there are still many things I haven't done, therefore I want to suggest several things in the application that can be added
 - add full funcionalitas
 - complete unit test
 - add log in application
 - should using transaction
 - mapping error database

## Thanks