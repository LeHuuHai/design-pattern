package main

type DBFactory interface {
	CreateConnection() Connection
	CreateQueryBuilder() QueryBuilder
	CreateTransaction() Transaction
}

type Connection interface {
	Connect()
}
type QueryBuilder interface {
	Build()
}
type Transaction interface {
	NewTransaction()
}
