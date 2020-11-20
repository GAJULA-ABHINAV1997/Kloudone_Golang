--At Publisher Port(9700)
CREATE DATABASE Abhinav;

CREATE TABLE person (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	email VARCHAR(150),
	gender VARCHAR(7) NOT NULL,
	date_of_birth DATE NOT NULL,
	country_of_birth VARCHAR(50) 
);

INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Dare', 'Mila', null, 'Male', '2020-01-25', 'China');
INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Prudi', 'Filimore', 'pfilimore1@reference.com', 'Female', '2020-01-09', 'Philippines');
INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Maxi', 'Swepson', 'mswepson2@i2i.jp', 'Female', '2020-04-10', 'South Africa');
INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Hilliard', 'Olligan', 'holligan3@opensource.org', 'Male', '2020-01-16', 'Philippines');
INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Katinka', 'Limpertz', 'klimpertz4@odnoklassniki.ru', 'Female', '2020-04-20', 'China');
INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Renato', 'Marner', 'rmarner5@latimes.com', 'Male', '2020-01-04', 'China');
INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Roxanne', 'Micklem', null, 'Female', '2019-12-11', 'Indonesia');

SELECT * FROM person;

CREATE PUBLICATION my_pub FOR ALL TABLES;

\dRp    -- Lists the publications



--At Subscriber port(9701)

CREATE TABLE person (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	email VARCHAR(150),
	gender VARCHAR(7) NOT NULL,
	date_of_birth DATE NOT NULL,
	country_of_birth VARCHAR(50) 
);

SELECT * FROM person;

CREATE SUBSCRIPTION my_sub CONNECTION 'host=127.0.01 port=9700 dbname=Abhinav' PUBLICATION my_pub;

\dRs    -- Lists the subscription




--At Publisher Port(9700)

INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES ('Gajue', 'Abhine', 'dhhdh7@gmal.com', 'Female', '2007-8-23', 'Sudan');


--At Subscriber port(9701)

SELECT * FROM person;

INSERT INTO person (id, first_name, last_name, email, gender, date_of_birth, country_of_birth) VALUES (79, 'Saia', 'cat', 'catg87@baby.com', 'Male', '2004-08-29', 'Germany');

SELECT * FROM person;



--At Publisher Port(9700)

SELECT * FROM person;