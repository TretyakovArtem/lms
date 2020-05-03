CREATE TABLE instance_client (
    instance_id integer REFERENCES instances(id) NOT NULL,
    client_id integer REFERENCES clients(id) NOT NULL,
    date_from varchar(100) NOT NULL,
    date_to varchar(100) NOT NULL,
    returned bool DEFAULT FALSE NOT NULL,
    return_date varchar(100) NOT NULL
);
