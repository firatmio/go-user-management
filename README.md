# User Management System in Go

![Go Version](https://img.shields.io/badge/Go-1.24.1-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Overview

This project is a simple yet complete user authentication system. It features a backend API built with Go (Golang) that handles user registration and login, and a basic frontend built with HTML, CSS, and JavaScript. User data is stored in a SQLite database.

## Features

-   **Secure User Registration:** Create new user accounts with a username and password.
-   **User Authentication:** Verify user credentials upon login.
-   **Password Hashing:** Passwords are securely hashed using the `bcrypt` algorithm.
-   **RESTful API:** A clean and simple API for user management.
-   **Simple Frontend:** A minimal web interface for registration and login.
-   **Database Persistence:** User data is stored in a local SQLite database file (`users.db`).

## Technologies Used

-   **Backend:** Go (Golang)
-   **Database:** SQLite
-   **Frontend:** HTML, CSS, JavaScript

## Getting Started

### Prerequisites

-   Go version 1.24.1 or higher. You can download it from [golang.org](https://golang.org/dl/).

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone <your-repository-url>
    cd user-management-go
    ```

2.  **Navigate to the backend directory:**
    ```bash
    cd backend
    ```

3.  **Run the server:**
    The Go module system will automatically handle the dependencies.
    ```bash
    go run .
    ```
    The server will start on `http://localhost:8080`.

4.  **Access the application:**
    Open your web browser and navigate to `http://localhost:8080`.

## API Endpoints

The backend server provides the following API endpoints:

| Method | Endpoint   | Description              |
| :----- | :--------- | :----------------------- |
| `POST` | `/register`| Registers a new user.    |
| `POST` | `/login`   | Authenticates a user.    |

### Request/Response Examples

#### `/register`

-   **Request Body:**
    ```json
    {
        "username": "testuser",
        "password": "password123"
    }
    ```
-   **Success Response (201 Created):** `User registered successfully`
-   **Error Response (409 Conflict):** `User already exists`

#### `/login`

-   **Request Body:**
    ```json
    {
        "username": "testuser",
        "password": "password123"
    }
    ```
-   **Success Response (200 OK):** `Login successful`
-   **Error Response (404 Not Found):** `User not found`
-   **Error Response (401 Unauthorized):** `Invalid password`

## Project Structure

```
.
├── backend/         # Go backend source code
│   ├── main.go      # Main application entry point, server setup
│   ├── handlers.go  # HTTP request handlers for API endpoints
│   ├── db.go        # Database initialization and connection
│   ├── models.go    # Struct definitions for data models
│   ├── go.mod       # Go module definitions
│   └── go.sum       # Go module checksums
├── frontend/        # Frontend files
│   ├── index.html   # Main HTML page
│   ├── style.css    # CSS for styling
│   └── script.js    # JavaScript for frontend logic
├── LICENSE          # Project License file
└── README.md        # This file
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
