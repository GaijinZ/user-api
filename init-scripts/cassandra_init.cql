CREATE KEYSPACE IF NOT EXISTS userapi WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};

CREATE TABLE IF NOT EXISTS userapi.users (id int PRIMARY KEY, firstname varchar, lastname varchar, email varchar);