package repository

const (
	queryDataRecovery = `SELECT JSON_AGG (rows)
		FROM( SELECT order_uid,
				track_number,
				entry,
				(SELECT TO_JSON (del)
					FROM (SELECT * FROM delivery WHERE deliveryId = id) AS del),
				(SELECT TO_JSON (pay)
					FROM (SELECT * FROM paymant WHERE paymantId = id) AS pay),
				(SELECT JSON_AGG (ite)
					FROM (SELECT * FROM items WHERE id = ANY(itemsId)) AS ite),
				locale,
				internal_signature,
				customer_id,
				delivery_service,
				shardkey,
				sm_id,
				date_created,
				oof_shard
			FROM order_us) AS rows;`
	queryGetOrder = `SELECT TO_JSON (rows)
		FROM( SELECT order_uid,
					track_number,
					entry,
					(SELECT TO_JSON (del)
						FROM (SELECT * FROM delivery WHERE deliveryId = id) AS del),
					(SELECT TO_JSON (pay)
						FROM (SELECT * FROM paymant WHERE paymantId = id) AS pay),
					(SELECT JSON_AGG (ite)
						FROM (SELECT * FROM items WHERE id = ANY(itemsId)) AS ite),
					locale,
					internal_signature,
					customer_id,
					delivery_service,
					shardkey,
					sm_id,
					date_created,
					oof_shard
				FROM order_us WHERE order_uid = $1) AS rows;`
	queryNewItems = `INSERT INTO items ( 
		chrt_id,
		track_number,
    	price,
    	rid,
    	name,
    	sale,
    	size,
    	total_price,
    	nm_id,
    	brand,
    	status)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id;`
	queryNewPaymant = `INSERT INTO paymant ( 
		transact,
    	request_id,
    	currency,
    	provider,
    	amount,
    	payment_dt,
    	bank,
    	delivery_cost,
    	goods_total,
    	custom_fee)
	VALUES($1, NULLIF($2, NULL), $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id;`
	queryNewDelivery = `INSERT INTO delivery ( 
		Name,
		Phone,
		Zip,
		City,
		Address,
		Region,
		Email)
	VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;`
	queryNewOrder = `INSERT INTO order_us ( 
		order_uid,
		track_number,
		entry,
		deliveryId,
		paymantId,
		itemsId,
		locale,
	    internal_signature,
	    customer_id,
	    delivery_service,
	    shardkey,
	    sm_id,
	    date_created,
	    oof_shard)
	VALUES($1, $2, $3, $4, $5, $6, $7, NULLIF($8, NULL), $9, $10, $11, $12, $13, $14);`

	queryCreateUser = `INSERT INTO user (
		uid,
    	name,
    	lastname,
    	number,
    	mail)
	VALUES($1, &2, &3, &4, &5);`

	queryCreateCommunication = `INSERT INTO communication (
		user1,
    	room)
	VALUES($1, &2);`

	queryCreateFriendRequest = `INSERT INTO friend_request (
		user1,
    	user2)
	VALUES($1, &2);`

	queryDeleteFriendRequest = `DELETE FROM friend_request 
	WHERE user1 = $1 AND 
	 user2 = $2;`

	queryCreateMessege = `INSERT INTO messege (
		author,
		recipient,
		data,
		time)
	VALUES($1, &2, &3, &4);`

	queryGetUsers = `SELECT uid, name, lastname FROM user;`

	queryGetFriends = `SELECT user2 FROM friend_request WHERE user1 = $1` //TODO JSON

	queryGetMesseges = `SELECT data, time FROM messege WHERE author = &1 AND recipient = $2` //todo поиск по uid комнаты

)