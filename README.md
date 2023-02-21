## Description Program

Design an API with the concept of REST using Go language with Echo Framework, and PostgreSQL as a Database.

## Database Design

![App Screenshot](https://github.com/abduromanov2020/ibj-technical-test/blob/main/database-design.png)

## API Contract & Documentation

https://documenter.getpostman.com/view/17874729/2s93CKPER6

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_URL`

    postgres://postgres:[db_password]@localhost:[db_port]/[db_name]

## How to Run Program

1. Start your PostgreSQL server

2. Clone the project

```bash
  git clone https://github.com/abduromanov2020/ibj-technical-test.git
```

3. Go to the project directory

```bash
  cd ibj-technical-test
```

4. Install dependencies

```bash
  go mod tidy
```

5. Start the server

```bash
  go run main.go
```

6. Use Postman collection to test it. Import the json into your Postman app, then send the message.

## Tech Stack

To run this app, you might to ensure youe machine has these instance installed:

1. Go Programming Language (1.18 or higher)
2. Postman or other RestAPI client
3. Code / Text Editor
4. PostgreSQL
