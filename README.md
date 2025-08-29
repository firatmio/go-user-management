# User Authentication System

A fully functional user registration and login system built with Go (Golang) backend, SQLite database, and a simple HTML/CSS/JavaScript frontend.

---

## Features

* Secure user registration (username & password)
* User login authentication
* Passwords are securely hashed using bcrypt
* Lightweight web interface
* Persistent storage using SQLite

---

## Setup Instructions

1. Clone the repository:

```bash
git clone https://github.com/yourusername/user-auth-system.git
cd user-auth-system/backend
```

2. Install required Go package:

```bash
go get golang.org/x/crypto/bcrypt
```

3. Start the server:

```bash
go run main.go
```

4. Open your browser at:

```
http://localhost:8080
```

---

## Usage

* **Register**: Fill out the registration form and click the `Register` button.
* **Login**: Fill out the login form and click the `Login` button.

---

## Project Structure

```
backend/      # Go backend files (API, database, models)
frontend/     # HTML, CSS, and JavaScript frontend files
README.md     # Project documentation
```

---

## Screenshots

*Add screenshots of the registration and login interface here to showcase the UI.*

---

## License

[MIT License](./LICENSE)