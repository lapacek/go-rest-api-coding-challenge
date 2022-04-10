CREATE SEQUENCE ticket_id_seq;

CREATE TABLE IF NOT EXISTS ticket (
    id integer NOT NULL DEFAULT nextval('ticket_id_seq'),
    first_name varchar(200) NOT NULL,
    last_name varchar(200) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    launchpad_id INT NOT NULL,
    destination_id INT NOT NULL,
    launch_date TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
    );

ALTER SEQUENCE ticket_id_seq OWNED BY ticket.id;

CREATE SEQUENCE launch_id_seq;

CREATE TABLE IF NOT EXISTS launch (
    id integer NOT NULL DEFAULT nextval('launch_id_seq'),
    first_name varchar(200) NOT NULL,
    last_name varchar(200) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    launchpad_id INT NOT NULL,
    destination_id INT NOT NULL,
    launch_date TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
    );

ALTER SEQUENCE launch_id_seq OWNED BY launch.id;