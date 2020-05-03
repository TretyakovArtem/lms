CREATE TABLE instance_publisher (
    instance_id integer REFERENCES instances(id) NOT NULL,
    publisher_id integer REFERENCES publishers(id) NOT NULL
);
