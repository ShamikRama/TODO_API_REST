#!/bin/bash

# Путь к папке с миграциями
MIGRATIONS_DIR="/Users/shamil/TODO_API_REST/migrations"

# Параметры подключения к базе данных
DB_USER="postgres"
DB_PASSWORD="mysecretpassword"
DB_NAME="shamil"

# Выполнение миграций
goose -dir $MIGRATIONS_DIR postgres "user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" up