# Exordium
**/ɪɡˈzɔːdɪəm/ - The Beginning of Anything.**

## Introduction
This repository includes all the tools you need to start your own industrial grade projects. This repository includes:
- Golang (Back End)
- ReactTS (Front End)
- RabbitMQ (Message Broker)
- PostgreSQL (Database)

The front end uses `yarn` insteam of the regular `npm` for further dependability. Install [yarn](https://classic.yarnpkg.com/lang/en/docs/install/) first.

Unlike [GDRP](https://github.com/acailuv/GDRP), this repository will only provide minimum demo since an advanced level knowledge of full web application development is expected from the user of this repository. this demo will only act as a boilerplate code for further development.

## Pre-requisites
This section will list out what you need to install before using this repository in your local development environment.
- Install [Go](https://golang.org/doc/install).
- Install [NodeJS](https://nodejs.org/en/download/).
- Install [Docker](https://docs.docker.com/get-docker/).
- Install [Docker Compose](https://docs.docker.com/compose/install/).
- Install [golang-migrate](https://github.com/golang-migrate/migrate).
    - You (might) need [Brew](https://brew.sh/).

## Basics
This section will cover basic usage of this repo.
- Starting Up Your Application:
    - `docker-compose up` will start your application.
- Makefile Macros:
    - `make genearte-migration name=add_user_table` will generate migration files.
    - `make run-migration-up database='postgres://root:your-secret-password@localhost:5432/app-db?sslmode=disable'` will run the migration (up).
    - `make run-migration-down database=''postgres://root:root@localhost:5432/app-db?sslmode=disable'` will run the migration (down).
    - **NOTE:** You might want to set your local database URL to your bash profile via `export` or something similar for extra convinience.

## Status
In development, to-do:
- Redis