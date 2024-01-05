# JournalAPI

## Introduction
JournalAPI is a RESTful API developed in Go, designed for managing journal entries. It utilizes Gin-Gonic for routing, GORM for object-relational mapping, Swagger for API documentation, and PostgreSQL for data storage. The application is containerized using Docker and orchestrated with Docker Compose.

## Features
- CRUD operations for users, journal entries, and journals.
- Dockerized environment for easy setup and deployment.
- Auto-generated Swagger documentation for easy API usage.
- PostgreSQL database integration.


## Prerequisites
- [Docker](https://www.docker.com)
- [Docker Compose](https://docs.docker.com/compose/install/)


## How 2 Run
**Clone the Repository**:
```bash
git clone https://github.com/borawhocodess/journalapi.git
```
**Start the Application**:  Navigate to the project directory and run:
```bash
docker-compose up --build
```

## Usage
Access the API through
- `http://localhost:8080/api/`
Swagger documentation can be found at
- `http://localhost:8080/swagger/index.html`


## Code Structure
- `main.go`: Application entry point.
- `/models`: Contains the User, Entry, and Journal models.
- `/database`: Database configuration and initialization.
- `/handlers`: Handler functions for API endpoints.


## How 2 Stop
**Stop the Application**:  Navigate to the project directory and run:
```bash
docker-compose down
```


