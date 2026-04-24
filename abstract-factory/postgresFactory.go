package main

import "fmt"

type PostgresConnection struct{}

func (_ *PostgresConnection) Connect() {
	fmt.Println("Postgres connected")
}

type PostgresQueryBuilder struct{}

func (_ *PostgresQueryBuilder) Build() {
	fmt.Println("Postgres query built")
}

type PostgresTransaction struct{}

func (_ *PostgresTransaction) NewTransaction() {
	fmt.Println("Postgres transaction begin")
}

type PostgresFactory struct{}

func (_ PostgresFactory) CreateConnection() Connection {
	return &PostgresConnection{}
}

func (_ PostgresFactory) CreateQueryBuilder() QueryBuilder {
	return &PostgresQueryBuilder{}
}

func (_ PostgresFactory) CreateTransaction() Transaction {
	return &PostgresTransaction{}
}
