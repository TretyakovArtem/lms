CREATE TABLE authors (
	id INT UNSIGNED auto_increment NOT NULL,
	name varchar(255) NOT NULL,
	year varchar(255) NOT NULL,
	CONSTRAINT authors_PK PRIMARY KEY (id),
	CONSTRAINT authors_UN UNIQUE KEY (id)
);