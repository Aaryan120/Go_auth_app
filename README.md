# Go Authentication System

This is a simple user authentication system built with **Go**, **MongoDB**, **Fiber** framework, and **Gmail SMTP** for OTP-based verification. The system provides the following functionalities:

- **User Registration**: Registers a new user with email and password.
- **OTP Verification**: Sends an OTP to the user's email for verification.
- **Login**: Allows a user to log in with email and password.

## Features

- **User Registration**: Users can sign up with an email and password. An OTP is sent to their email for verification.
- **Email OTP Verification**: The user must verify their email by entering the OTP sent to their inbox.
- **Secure Password Storage**: User passwords are hashed using **bcrypt** before being stored in the database.
- **Login**: Users can log in after they have verified their email address.
- **Environment Variables**: Uses environment variables for sensitive configurations such as email settings.

## Prerequisites

To run this project, you need the following installed on your machine:

- **Go**: To run the Go application. Install it from [here](https://golang.org/dl/).
- **MongoDB**: A running MongoDB instance to store user data. You can use a local MongoDB instance or a cloud-based solution like [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).
- **Gmail Account**: For sending OTP emails. You need a Gmail account and app-specific password (if using 2-step verification).

## Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/go-auth-app.git
   cd go-auth-app
