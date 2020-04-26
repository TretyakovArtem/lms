CREATE TABLE instance_client (
	instance_id INT UNSIGNED NOT NULL,
	client_id INT UNSIGNED NOT NULL,
	date_from varchar(255) NOT NULL,
	date_to varchar(255) NOT NULL,
	returned BOOL DEFAULT 0 NOT NULL,
	return_date varchar(255) NULL,
	CONSTRAINT instance_client_PK PRIMARY KEY (instance_id,client_id,date_from),
	CONSTRAINT instance_client_UN UNIQUE KEY (instance_id,client_id,date_from),
	CONSTRAINT instance_client_FK FOREIGN KEY (instance_id) REFERENCES instances(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT instance_client_FK_1 FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE ON UPDATE CASCADE
);