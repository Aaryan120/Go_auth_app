# Go Authentication System

This is a simple user authentication system built with **Go**, **MongoDB** and **Fiber** framework. The system provides the following functionalities:

- **User Registration**: Registers a new user with email and password.
- **Login**: Allows a user to log in with email and password.

## Features

- **User Registration**: Users can sign up with an email and password.
- **Secure Password Storage**: User passwords are hashed using **bcrypt** before being stored in the database.
- **Login**: Users can log in after they have verified their email address.
- **Environment Variables**: Uses environment variables for sensitive configurations such as email settings.

## Prerequisites

To run this project, you need the following installed on your machine:

- **Go**: To run the Go application. Install it from [here](https://golang.org/dl/).
- **MongoDB**: A running MongoDB instance to store user data. You can use a local MongoDB instance or a cloud-based solution like [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

## Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Aaryan120/go-auth-app.git
   cd go-auth-app
