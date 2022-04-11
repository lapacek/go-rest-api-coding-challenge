package data

const INSERT_TICKET string = "INSERT INTO ticket(launch_id, first_name, last_name, gender, birthday) " +
								"VALUES($1, $2, $3, $4, $5)"

const INSERT_LAUNCH string = "INSERT INTO launch(launchpad_id, destination_id, launch_date) " +
								"VALUES($1, $2, $3)"

const SELECT_LAUNCHES string = "SELECT id, launchpad_id, destination_id, launch_date FROM launch " +
								" WHERE launch_date BETWEEN $1 AND $2"

const SELECT_LAUNCH string = "SELECT id, launchpad_id, destination_id, launch_date FROM launch " +
								" WHERE launch_date = $1"

const SELECT_DESTINATIONS string = "SELECT id, name FROM destination"

const SELECT_LAUNCHPADS string = "SELECT id, id_spacex FROM launchpads"

const SELECT_BOOKINGS string = "SELECT t.first_name, t.last_name, t.gender, t.birthday, " +
								"l.launchpad_id, l.destination_id, l.launch_date" +
								"FROM ticket t " +
								"JOIN launch l ON t.launch_id = l.id"

