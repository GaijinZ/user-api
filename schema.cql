CREATE KEYSPACE IF NOT EXISTS usersapi WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};

CREATE TABLE IF NOT EXISTS usersapi.users (id int PRIMARY KEY, firstname varchar, lastname varchar, email varchar);