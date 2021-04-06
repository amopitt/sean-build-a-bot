CREATE TABLE cart (
	cart_id serial PRIMARY KEY,
	robot_name VARCHAR ( 50 )  NOT NULL,
	cost MONEY NOT NULL, 
	created_on TIMESTAMP NOT NULL 
);