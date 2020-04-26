CREATE TABLE book_instance (
	book_id INT UNSIGNED NOT NULL,
	instance_id INT UNSIGNED NOT NULL,
	CONSTRAINT book_instance_UN UNIQUE KEY (book_id,instance_id),
	CONSTRAINT book_instance_FK FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT book_instance_FK_1 FOREIGN KEY (instance_id) REFERENCES instances(id) ON DELETE CASCADE ON UPDATE CASCADE
);
