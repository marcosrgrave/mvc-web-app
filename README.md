# MVC Web Application

This project is a study to apply Golang in web applications.
The goal is to come up with a simple CRUD application.
Golang was used for the backend, HTML and CSS for the frontend and MySQL as database.

## How to run it

1. Golang and Docker required.

2. At the project root directory, run:

   ```bash
   docker-compose up -d
   ```

   This will initialize the MySQL database.

3. Wait 10 seconds and then run:

   ```bash
   go run main.go
   ```

4. Access the local server in your browser:

   ```txt
   http://localhost:8080/
   ```

5. Have fun =]

6. After that, press `ctrl + c` in the terminal to stop de backend and `docker-compose down` to stop de database.

## Project's Goals

- [X] Docker Compose (Services and Database)
- [X] Multi-Stage Build (greatly reduces Docker image size)
- [x] CRUD
- [x] Models DTO Conversion
- [x] Database Connection (MySQL)
- [x] MVC Structure
- [x] HTML and CSS Frontend
  - [x] Home Page (Categories and Products)
  - [X] List of All Products Page + Update button + Delete button
  - [x] New Product Page
  - [X] Update Existing Product Page
  - [ ] List of All Categories Page + Update button + Delete button
  - [ ] New Category Page
  - [ ] Update Existing Category Page
- [ ] Environment Variables
- [ ] Tests
- [ ] Code documentation (Functions and Methods)
- [ ] API documentation with Swagger and OpenAPI 3.0
- [ ] Authentication
