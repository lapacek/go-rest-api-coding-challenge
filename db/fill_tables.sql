INSERT INTO launchpad (id_spacex, name)
VALUES ('5e9e4501f5090910d4566f83', 'VAFB SLC 3W'),
       ('5e9e4501f509094ba4566f84', 'CCSFS SLC 40'),
       ('5e9e4502f5090927f8566f85', 'STLS'),
       ('5e9e4502f5090995de566f86', 'Kwajalein Atoll'),
       ('5e9e4502f509092b78566f87', 'VAFB SLC 4E'),
       ('5e9e4502f509094188566f88', 'KSC LC 39A');

INSERT INTO destination (name)
VALUES ('Mars'),
       ('Moon'),
       ('Pluto'),
       ('Asteroid Belt'),
       ('Europa'),
       ('Titan'),
       ('Ganymede');

INSERT INTO launch (id, launchpad_id, destination_id, launch_date)
VALUES ('1', '1', '1', '2020-11-11'),
       ('2', '2', '3', 'Rimmer', '2022-01-01');

INSERT INTO ticket (id, launch_id, first_name, last_name, gender, birthday)
VALUES ('1', '1', 'Chuck', 'Norris', 'men', '1940-03-10'),
       ('2', '2', 'Arnold', 'Rimmer', 'hologram', '1989-01-01');
