package repository

const (
	queryCreateUser = `INSERT INTO users (
		uid,
    	name,
    	lastname,
    	number,
    	mail)
	VALUES($1, $2, $3, $4, $5);`

	queryCreateCommunication = `INSERT INTO communication (
		user1,
    	room)
	VALUES($1, $2);`

	queryCreateFriendRequest = `INSERT INTO friend_request (
		user1,
    	user2)
	VALUES($1, $2);`

	queryDeleteFriendRequest = `DELETE FROM friend_request 
	WHERE user1 = $1 AND 
	 user2 = $2;`

	queryCreateMessege = `INSERT INTO messege (
		author,
		recipient,
		data,
		time)
	VALUES($1, $2, $3, $4);`

	queryGetUsers = `SELECT JSON_AGG (rows) FROM(SELECT uid, name, lastname FROM users) AS rows;`

	queryGetFriends = `SELECT JSON_AGG (rows) FROM(
		SELECT uid, name, lastname 
		FROM users
		WHERE uid = ANY(
			SELECT user1 
			FROM communication 
			WHERE room = ANY(
				SELECT room 
				FROM communication 
				WHERE user1 = $1) 
			AND user1 != $1
		)
	) AS rows;` //TODO JSON

	queryGetMesseges = `SELECT JSON_AGG (rows) FROM(
		SELECT 	author,
				data, 
				time 
		FROM messege 
		WHERE (author = $1 AND recipient = $2) OR (author = $2 AND recipient = $1)) AS rows;` //todo поиск по uid комнаты

	queryGetFriendRequest = `SELECT JSON_AGG (rows) FROM(
			SELECT uid, name, lastname
			FROM users
			WHERE uid = ANY (
				SELECT user1 
				FROM friend_request 
				WHERE user2 = $1)
	) AS rows;`

	queryGetKey = `SELECT JSON_AGG (rows) FROM(
			SELECT data 
			FROM messege 
			WHERE time = (SELECT MIN(time) FROM messege) AND author = $1 AND recipient = $2
		) AS rows;`

	queryDeleteKey = `DELETE FROM messege WHERE time = (SELECT MIN(time) FROM messege) AND author = $1 AND recipient = $2;`
)