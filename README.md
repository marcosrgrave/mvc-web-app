# MVC Web Application 

Note: this is still a work in progress.

This project was created to study Golang applied to web applications.
The goal is to come up with a robust application, capable of simulating a simple but real scenario, and highly scalable.
Golang was used for the backend, HTML and CSS for the frontend and MySQL for the database.

## How to run this web application

1. Golang and Docker required.

2. At the project root directory, run:
   ```
   docker-compose up -d
   ```
   This will initialize the MySQL database.

3. Wait 10 seconds and then run this command:
   ```
   go run main.go
   ```

4. Access the local server in your browser:
   ```
   http://localhost:8080/
   ```
5. Have fun =]

6. After that, press "ctrl + c" in the terminal to stop de backend and "docker-compose down" to stop de database.

## Project's Goals

- [ ] Microservices
  - [X] Docker Compose (Services and Database)
    - [X] APIs with Multi-Stage Build (greatly reduces Docker image size)
    - [ ] APIs Environment Variables
  - [ ] API Gateway with Kong and Konga
- [ ] Code documentation (Functions and Methods)
- [ ] API documentation with Swagger and OpenAPI 3.0
- [x] CRUD
- [x] MVC Structure
- [x] HTML and CSS Frontend
  - [x] Home Page (Categories and Products)
  - [ ] List of All Products Page + Update button + Delete button
  - [x] New Product Page
  - [ ] Update Existing Product Page
  - [ ] List of All Categories Page + Update button + Delete button
  - [ ] New Category Page
  - [ ] Update Existing Category Page
- [x] Models DTO Conversion
- [x] Database Connection (MySQL)
- [ ] Tests
- [ ] Authentication
- [ ] Ports and Adapters