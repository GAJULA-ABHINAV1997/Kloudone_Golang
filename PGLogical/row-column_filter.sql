--Publication port:9700

CREATE DATABASE students;

CREATE EXTENSION pglogical;

CREATE TABLE students(id INT PRIMARY KEY,name TEXT,address TEXT,city TEXT);

SELECT pglogical.create_node(
    node_name := 'new_prov',
    dsn := 'host=localhost port=9700 dbname=students'
    );

SELECT pglogical.replication_set_add_table(
    set_name := 'default', relation :='students', 
    synchronize_data := true, 
    columns :='{id,name,city}', 
    row_filter := 'id<5'
    );

INSERT INTO students VALUES(1,'sai',' Nagar','Vij');

INSERT INTO students VALUES(2,'Kun','Mount','rjy');

INSERT INTO students VALUES(3,'jai','Anna street','Tuticorion');

INSERT INTO students VALUES(4,'Sunnanda','jp nagar','vjy');

INSERT INTO students VALUES(5,'Pavan','kalmkuko','kahmmam');

SELECT * FROM students;

--  Subscriberport:9701

CREATE EXTENSION Pglogical;

CREATE TABLE students(id INT PRIMARY KEY,name TEXT,address TEXT,city TEXT);

SELECT pglogical.create_node(
    node_name := 'new_sub',
    dsn := 'host=localhost port=9701 dbname=students'
    );

SELECT pglogical.create_subscription( 
    subscription_name := 'new_subp',
    provider_dsn := 'host=localhost port=9700 dbname=students'
    );

SELECT * FROM students;