CREATE SEQUENCE destination_id_seq;

CREATE TABLE IF NOT EXISTS destination (
    id integer NOT NULL DEFAULT nextval('destination_id_seq'),
    name varchar(200) NOT NULL,
    PRIMARY KEY (id)
    );

ALTER SEQUENCE destination_id_seq OWNED BY destination.id;;

CREATE SEQUENCE launchpad_id_seq;

CREATE TABLE IF NOT EXISTS launchpad (
    id integer NOT NULL DEFAULT nextval('launchpad_id_seq'),
    id_spacex varchar(200) NOT NULL,
    name varchar(200) NOT NULL,
    PRIMARY KEY (id)
    );

ALTER SEQUENCE launchpad_id_seq OWNED BY launchpad.id;

CREATE SEQUENCE launch_id_seq;

CREATE TABLE IF NOT EXISTS launch (
    id integer UNIQUE NOT NULL DEFAULT nextval('launch_id_seq'),
    launchpad_id INT NOT NULL,
    destination_id INT NOT NULL,
    launch_date DATE NOT NULL,
    PRIMARY KEY (launchpad_id, launch_date),
    CONSTRAINT fk_launchpad
        FOREIGN KEY(launchpad_id)
            REFERENCES launchpad(id),
    CONSTRAINT fk_destination
        FOREIGN KEY(destination_id)
            REFERENCES destination(id)
    );

ALTER SEQUENCE launch_id_seq OWNED BY launch.id;

CREATE SEQUENCE ticket_id_seq;

CREATE TABLE IF NOT EXISTS ticket (
    id integer NOT NULL DEFAULT nextval('ticket_id_seq'),
    launch_id INT NOT NULL,
    first_name varchar(200) NOT NULL,
    last_name varchar(200) NOT NULL,
    gender varchar(200) NOT NULL,
    birthday DATE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_launch
        FOREIGN KEY(launch_id)
            REFERENCES launch(id)
    );

ALTER SEQUENCE ticket_id_seq OWNED BY ticket.id;


