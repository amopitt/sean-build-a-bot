CREATE TABLE orders (
	order_id serial PRIMARY KEY,
	robot_name VARCHAR ( 50 )  NOT NULL,
	cost NUMERIC(5,2) NOT NULL, 
	created_on TIMESTAMP NOT NULL 
);