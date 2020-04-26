CREATE TABLE producers (
	id INT UNSIGNED auto_increment NOT NULL,
	name varchar(255) NOT NULL,
	CONSTRAINT producers_PK PRIMARY KEY (id),
	CONSTRAINT producers_UN UNIQUE KEY (id)
);