package sql

const InsertUser = "INSERT INTO users (first_name, last_name, birth_date) VALUES ($1,$2,$3) RETURNING id"
const InsertAddress = "INSERT INTO addresses (user_id, street, number, city) VALUES ($1,$2,$3,$4) RETURNING id"

const GetUserById = "SELECT u.id, u.first_name, u.last_name, u.birth_date FROM users u WHERE id=$1"
const GetAddressesByUserId = "SELECT a.id, a.user_id, a.street, a.number, a.city FROM addresses a WHERE user_id=$1"

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
	WHERE u.id = ANY ($1)
`
