# BlogPost API

A simple RESTful API for a blogging platform built in Go and using PostgreSQL as the database.

![License](https://img.shields.io/badge/license-MIT-blue)

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Database](#database)
- [Contributing](#contributing)
- [License](#license)

## Description

The BlogPost API is a straightforward RESTful API designed to manage blog posts. It provides basic CRUD (Create, Read, Update, Delete) functionality for blog posts, stored in a PostgreSQL database. The API is built in Go and follows RESTful conventions.

## Features

- Create, retrieve, update, and delete blog posts.
- Store blog posts in a PostgreSQL database.
- API endpoints for listing all blog posts and fetching a specific post by ID.
- Proper error handling and response status codes.
- Request logging and CORS support.

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL installed and configured.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/berbreik/BlogPost.git
2.
   Set environment variables for your PostgreSQL connection:
  ```bash
  export DB_CONNECTION_STRING="postgresql://user:password@localhost/blogpost?sslmode=disable"
  export PORT="8080"
3. Initialize and run the application:
   ```bash
   go run main.go

API Endpoints :

Retrieve a list of all blog posts
URL: /v1/posts
Method: GET

Retrieve a specific blog post by ID
URL: /v1/posts/{id}
Method: GET

Create a new blog post
URL: /v1/posts
Method: POST

Update an existing blog post
URL: /v1/posts/{id}
Method: PUT

Delete a blog post
URL: /v1/posts/{id}
Method: DELETE
