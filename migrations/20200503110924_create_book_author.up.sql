CREATE TABLE book_author (
    book_id integer REFERENCES books(id) NOT NULL,
    author_id integer REFERENCES authors(id) NOT NULL
);
