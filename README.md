# DBO Technical Test Project

Welcome to the DBO technical test project repository. This project demonstrates the implementation of multiple services including `customer_service`, `order_service`, and `auth_service`, alongside Docker configurations and database migrations.

## Table of Contents

- [Project Overview](#project-overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Database Migration](#database-migration)
  - [Running with Docker](#running-with-docker)
  - [Running Locally](#running-locally)
- [License](#license)

## Project Overview

This project is designed as a technical test for DBO to showcase proficiency in Go programming, Docker, and database management. It includes essential services and migration scripts for a functional application setup.

## Getting Started

### Prerequisites

Before running the project, ensure you have the following installed:

- Go 1.21.6 or later
- Docker (for Docker-based setup)
- `golang-migrate` for database migrations

### Database Migration

To set up the database schema, you need to run the migration scripts. Follow these steps:

1. **Install `golang-migrate`** if not already installed. You can download it from the [official site](https://github.com/golang-migrate/migrate).

2. **Apply the Migrations**

   Run the following command to apply migrations:

   ```sh
   migrate -path ./migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" up
