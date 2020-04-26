CREATE TABLE books (
	id INT UNSIGNED auto_increment NOT NULL,
	name varchar(255) NOT NULL,
	year varchar(255) NOT NULL,
	CONSTRAINT books_PK PRIMARY KEY (id),
	CONSTRAINT books_UN UNIQUE KEY (id)
);