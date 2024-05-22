# GO REST API

This is a simple REST API built with Go, using the Gin-Gonic framework. It includes a basic login system implemented with JWT authentication, allowing users to register, login, and manage sessions securely. It also supports event management where authenticated users can create, manage, and register for events.

## Features

- **User Registration and Authentication**: Securely register and authenticate users using JWT.
- **Event Management**: Authenticated users can create, edit, and delete events. Users can also register for events and cancel their registrations.

## Tech Stack

- **Go**: Primary programming language.
- **Gin-Gonic**: HTTP web framework used to build the API.
- **SQLite**: Database used for storing user and event data.

## Installation

To get this project up and running on your machine, follow these steps:

1. Clone the repository:

git clone [<repository-url>](https://github.com/Mlcarvalho1/go-rest-api.git)

2. Navigate to the project directory:

cd go-rest-api

3. Install the dependencies:

go mod tidy

4. Run the application:

go run main.go


## API Endpoints

### User Endpoints

- **POST /signup**: Register a new user.
- **POST /login**: Login for existing users.

### Event Endpoints

- **POST /events**: Create a new event.
- **GET /events**: Retrieve all events.
- **GET /events/:id**: Get details of a specific event.
- **PUT /events/:id**: Update a specific event.
- **DELETE /events/:id**: Delete a specific event.
- **POST /events/:id/register**: Register for an event.
- **DELETE /events/:id/register**: Cancel event registration.

## Configuration

Before running the application, ensure you have set up the necessary environment variables, including the secret key for JWT generation.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.


