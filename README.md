# online-store-putridefa

## Description

This project is a test challenge for becoming a Backend Engineer at PT. Synapsis Sinergi Digital.

## Table of Contents

- [Installation](#installation)
- [Migration](#migration)
- [Running](#running)

## Installation


### Prerequisites

- Go
- PostgreSql

### Installation Steps

# Clone the repository
git clone https://github.com/yantidefa/online-store.git

# Navigate to the project directory
cd online-store

# Install dependencies
go get -u -v

## Migration

### Navigate to the migrations directory
cd migrations

### Status Migration
goose postgres "user=postgres dbname=online_store host=localhost port=5432 password=0303 sslmode=disable" status

### Up Migration
goose postgres "user=postgres dbname=online_store host=localhost port=5432 password=0303 sslmode=disable" up

### Down Migration
goose postgres "user=postgres dbname=online_store host=localhost port=5432 sslmode=disable" down

# Running the project
go run main.go
