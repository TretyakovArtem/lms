CREATE TABLE instances (
	id INTEGER UNSIGNED auto_increment NOT NULL,
	producer_id INT UNSIGNED NOT NULL,
	`year` varchar(255) NOT NULL,
	CONSTRAINT instances_PK PRIMARY KEY (id),
	CONSTRAINT instances_UN UNIQUE KEY (id),
	CONSTRAINT instances_FK_1 FOREIGN KEY (producer_id) REFERENCES producers(id) ON DELETE CASCADE ON UPDATE CASCADE
);