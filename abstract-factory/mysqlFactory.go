package main

import "fmt"

type MySQLFactory struct{}

type MySQLConnection struct{}

func (_ *MySQLConnection) Connect() {
	fmt.Println("MySQL connected")
}

type MySQLQueryBuilder struct{}

func (_ *MySQLQueryBuilder) Build() {
	fmt.Println("MySQL query built")
}

type MySQLTransaction struct{}

func (_ *MySQLTransaction) NewTransaction() {
	fmt.Println("MySQL transaction begin")
}
func (_ MySQLFactory) CreateConnection() Connection {
	return &MySQLConnection{}
}

func (_ MySQLFactory) CreateQueryBuilder() QueryBuilder {
	return &MySQLQueryBuilder{}
}

func (_ MySQLFactory) CreateTransaction() Transaction {
	return &MySQLTransaction{}
}
