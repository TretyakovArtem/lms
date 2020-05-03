CREATE TABLE instances (
	id serial NOT NULL PRIMARY KEY,
	"year" varchar(255) NOT NULL,
    book_id integer REFERENCES books(id) NOT NULL,
    publisher_id integer REFERENCES publishers(id) NOT NULL
);
