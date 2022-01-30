package sql

const InsertUser = "INSERT INTO users (first_name, last_name, birth_date) VALUES ($1,$2,$3) RETURNING id"

const InsertAddresses = "INSERT INTO addresses (user_id, street, number, city) VALUES (:user_id, :street, :number, :city)"

const InsertAddress = "INSERT INTO addresses (user_id, street, number, city) VALUES ($1, $2, $3, $4) RETURNING id"

const GetUserWithAddressesById = `
	SELECT 
		u.id, 
		u.first_name, 
		u.last_name, 
		u.birth_date,
		a.id, 
		a.user_id, 
		a.street, 
		a.number, 
		a.city
	FROM users u
	LEFT JOIN addresses a ON u.id = a.user_id
	WHERE u.id=$1
`

const GetUsersWithAddressesByIds = `
	SELECT 
		u.id, 
		u.first_name, 
		u.last_name, 
		u.birth_date,
		a.id, 
		a.user_id, 
		a.street, 
		a.number, 
		a.city
	FROM users u
	LEFT JOIN addresses a ON u.id = a.user_id
	WHERE u.id IN (?)
`

const UpdateUser = "UPDATE users SET first_name=$1, last_name=$2, birth_date=$3 WHERE id=$4"

const DeleteUser = "DELETE FROM users WHERE id = $1"

const DeleteAddress = "DELETE FROM addresses WHERE id = $1"

const GetAddressById = `
	SELECT
		a.id, 
		a.user_id, 
		a.street, 
		a.number, 
		a.city
	FROM addresses a
	WHERE a.id=$1
`
