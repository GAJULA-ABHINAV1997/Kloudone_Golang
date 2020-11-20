-----------------------------------------------------------------
----------  PGLOGICAL WITH 1 PROVIDER IN 2 SUBSCRIBER  ----------
-----------------------------------------------------------------

-- PROVIDER (Port-9700)

CREATE EXTENSION pglogical;

CREATE TABLE tbl (id int primary key,name varchar);

SELECT pglogical.create_node (
    node_name := 'provider',
    dsn := 'host=localhost port=9700 dbname=Abhinav'
);

SELECT pglogical.replication_set_add_table (
    set_name := 'default', relation := 'tbl', synchronize_data := true
);



-- SUBSCRIBER1 (Port-9701 ,DB-sub1)

CREATE EXTENSION pglogical;

CREATE TABLE tbl (id int primary key,name varchar);

SELECT pglogical.create_node (
    node_name := 'subscriber1',
    dsn := 'host=localhost port=9701 dbname=sub1'
);

SELECT pglogical.create_subscription (
    subscription_name := 'subscription1',
    provider_dsn := 'host=localhost port=9700 dbname=Abhinav'
);



-- SUBSCRIBER2 (Port-9701 DB-sub2)

CREATE EXTENSION pglogical;

CREATE TABLE tbl (id int PRIMARY KEY, name varchar);

SELECT pglogical.create_node (
    node_name := 'subscriber2',\c
    dsn := 'host=localhost port=9701 dbname=sub3'
);

SELECT pglogical.create_subscription (
    subscription_name := 'subscribtion2',
    provider_dsn := 'host=localhost port=9700 dbname=Abhinav'
);


-- PROVIDER (Port-9700)

INSERT INTO tbl VALUES (1, 'Gajula');
INSERT INTO tbl VALUES (2, 'Sai');
INSERT INTO tbl VALUES (3, 'Manas');
INSERT INTO tbl VALUES (4, 'Jagan');

SELECT * FROM tbl;



-- SUBSCRIBER1 (Port-9701,DB-sub1)

SELECT * FROM tbl;



-- SUBSCRIBER2 (Port-9701,DB-sub2)

SELECT * FROM tbl;