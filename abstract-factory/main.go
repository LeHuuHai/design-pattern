package main

func run(factory DBFactory) {
	conn := factory.CreateConnection()
	qb := factory.CreateQueryBuilder()
	tx := factory.CreateTransaction()

	conn.Connect()
	tx.NewTransaction()
	qb.Build()
}

func main() {
	run(PostgresFactory{})
	run(MySQLFactory{})
}
