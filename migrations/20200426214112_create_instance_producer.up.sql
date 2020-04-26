CREATE TABLE instance_producer (
	producer_id INT UNSIGNED NOT NULL,
	instance_id INT UNSIGNED NOT NULL,
	CONSTRAINT instance_producer_UN UNIQUE KEY (instance_id,producer_id),
	CONSTRAINT instance_producer_FK FOREIGN KEY (producer_id) REFERENCES producers(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT instance_producer_FK_1 FOREIGN KEY (instance_id) REFERENCES instances(id) ON DELETE CASCADE ON UPDATE CASCADE
);