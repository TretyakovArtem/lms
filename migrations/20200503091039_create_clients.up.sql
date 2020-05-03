CREATE TABLE clients (
	id serial NOT NULL PRIMARY KEY,
	"name" varchar(255) NOT NULL,
    "phone" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL,
    "address" varchar(255) NOT NULL
);
