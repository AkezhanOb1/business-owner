package config

import "fmt"

//ServerAddress is
var ServerAddress = "0.0.0.0:50056"
//postgresAddress is the address of the postgres
//const postgresAddress = "127.0.0.1"
const postgresAddress = "46.101.138.224"

//postgresPort is the port of the postgres
const postgresPort = "5432"

//postgresDataBase is the name of the database
const postgresDataBase = "diploma"

//postgresUsername is the name of the user inside DBA
const postgresUsername = "postgres"

//postgresPassword is the password of the user
const postgresPassword = "postgres"

//PostgresConnection is the connection string to the database
var PostgresConnection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	postgresAddress, postgresPort, postgresUsername, postgresPassword, postgresDataBase)


//rabbitUserName is the name of the user in Rabbit MQ
const rabbitUserName = "admin"
//rabbitPassword is the password of the user
const rabbitPassword = "admin"
//RabbitConnection is the connection string to the Rabbit MQ
var RabbitConnection = fmt.Sprintf("amqp://%s:%s@46.101.138.224:5672/", rabbitUserName, rabbitPassword)
//WelcomeEmilQueue is
var WelcomeEmilQueue = "welcome"
//ResetPasswordQueue
var ResetPasswordQueue = "reset"