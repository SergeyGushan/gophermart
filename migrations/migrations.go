package migrations

func CreateUsersTableSql() string {
	return `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, 
		login VARCHAR(255) UNIQUE NOT NULL, 
		password VARCHAR(255) NOT NULL, 
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 	);`
}

func CreateOrdersTableSql() string {
	return `CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY, order_id VARCHAR(255) UNIQUE NOT NULL, 
		user_id INTEGER REFERENCES users(id) NOT NULL, 
		status varchar NOT NULL, 
		accrual float DEFAULT 0, 
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
}

func CreateOperationsTableSql() string {
	return `CREATE TABLE IF NOT EXISTS operations (
		id SERIAL PRIMARY KEY, order_id VARCHAR(255)  NOT NULL, 
		user_id INTEGER REFERENCES users(id) NOT NULL, 
		sum float NOT NULL, 
		type varchar  NOT NULL, 
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
}
