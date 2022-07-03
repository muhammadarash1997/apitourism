# API Tourism
> This is a backend service for API Tourism. API Tourism is used by travelers across Indonesia to search destinations along with complete information
such as profile, image, type, rating, popularity, and distance to the user current location. User can bookmark a destination, give rating to a destination,
search destinations, search nearby destinations by type, rating or popularity. The system used for knowing location is Geohash method.

## Usage
Create .env file in the root directory of your project. Add environment-specific variables on new lines in the form NAME=VALUE. For example:

```bash
- DB_HOST=localhost
- DB_PORT=5432
- DB_USER=root
- DB_PASS=password123
- DB_NAME=attendancedb
- TOKEN_HOUR_LIFESPAN=24
- SECRET_KEY=secret123
```

## Requirements
You need [PostgreSQL](https://www.postgresql.org/) = `14` server running.

## Running App

```bash
go run main.go
```

## Technologies Used
- Go 1.18
- PostgreSQL 14
- Gin Web Framework
- GORM
- Swagger 2.0
- JSON Web Token

## Usecases
1. There is service to register and login.
2. There is service to bookmark destination.
3. There is service to add the number of destination views.
3. There is service to give a rate to destination.
4. There is service to get nearby destinations by type, rating or popularity.
5. There is service to get all destinations.

## Code Structure
The design contains several layers and components and very much similar to onion or clean architecture attempt.

### Components
1. Controllers
2. Services
3. Repositories

#### Controllers
Controllers is where all the http handlers exist. This layer is responsible to hold all the http handlers and request validation.

#### Services
Services mediates communication between a controller and repository layer. The service layer contains business logic.

#### Repositories
Repositories is for accessing the database and helps to extend the CRUD operations on the database.
