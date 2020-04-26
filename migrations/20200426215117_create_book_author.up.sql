CREATE TABLE book_author (
	book_id INT UNSIGNED NOT NULL,
	author_id INT UNSIGNED NOT NULL,
	CONSTRAINT book_author_PK PRIMARY KEY (book_id,author_id),
	CONSTRAINT book_author_UN UNIQUE KEY (book_id,author_id),
	CONSTRAINT book_author_FK FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT book_author_FK_1 FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE ON UPDATE CASCADE
);