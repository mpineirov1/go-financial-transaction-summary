## Description

This project involves creating a system that processes a file containing debit and credit transactions on an account and sends summarized information to a user via email. The transaction file will be in CSV format and will include transactions with debit (-) and credit (+) indicators. The system will calculate the total balance in the account, the number of transactions grouped by month, and the average debit and credit amounts grouped by month.

## Requirements

- go 1.20
- Postgresql

## Configuration

1. In the root folder, create a file named .env if it does not already exist. Within this file, set the following properties to configure your mail settings:

```
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_USER=your_db_usr
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

MAIL_HOST=your_mail_host
MAIL_PORT=your_mail_port
MAIL_FROM_ADDRESS=your_from_address
MAIL_USERNAME=your_mail_username
MAIL_PASSWORD=your_mail_password

TEST_EMAIL=your_test_email
```

## Running 

1. Execute go run main.go
