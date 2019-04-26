# Go Starter App

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

This is a basic CRUD using GOLANG

You can access the API in: https://go-apis.herokuapp.com

## Database (MYSQL)
- Create an database with name "school" and one table like this (MYSQL example):
```
CREATE TABLE `students` (
  `id` bigint(200) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `course` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

```

## Install
- You can install GO checking the documentation: https://golang.org/doc/install
- After install the GO and cloning the repository, access the root directory  and execute: go run ./src/api/main.go 

## Routes
- GET /students = All students from database

```
--> response
[
    {
        "id": 1,
        "name": "otavio rojas",
        "age": 23,
        "email": "golang@gmail.com",
        "course": "Ciencia da computacao"
    },
    {
        "id": 7,
        "name": "Rodolfo",
        "age": 28,
        "email": "rodolfo@gmail.com",
        "course": "Computer Science"
    },
    {
        "id": 10,
        "name": "Ramalho",
        "age": 45,
        "email": "golangjr@gmail.com",
        "course": "Medicine"
    }
]
```

- GET /student/:id = Search an student with id

```
--> response
{
    "name": "Golang Jr",
    "age": 20,
    "email": "golang@gmail.com",
    "course": "Computer Science"
}
```

- POST /student = Save an student in database

```
--> request
{
    "name": "Golang Jr",
    "age": 20,
    "email": "golang@gmail.com",
    "course": "Computer Science"
}
```
```
--> response
{
    "name": "Golang Jr",
    "age": 20,
    "id": 40
    "email": "golang@gmail.com",
    "course": "Computer Science"
}
```

- DELETE /student/:id = Delete an student with id

```
--> response
{
    "message": "Student with id X was deleted"
}

```
- PUT /student/:id = Update an student with id

```
--> request
{
    "age": 22,
}
```

```
--> response
{
    "name": "Golang Jr",
    "age": 22,
    "id": 40
    "email": "golang@gmail.com",
    "course": "Computer Science"
}
```

4. Start coding!

## Questions

* [Otavio Rojas](xrojasinc@gmail.com)
