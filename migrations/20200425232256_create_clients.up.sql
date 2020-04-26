CREATE TABLE clients (
	id INT UNSIGNED auto_increment NOT NULL,
	name varchar(255) NOT NULL,
	address varchar(255) NOT NULL,
	phone varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	CONSTRAINT clients_PK PRIMARY KEY (id),
	CONSTRAINT clients_UN UNIQUE KEY (id)
);