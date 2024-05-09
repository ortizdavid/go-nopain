DROP DATABASE IF EXISTS db_api_key;
CREATE DATABASE db_api_key;

\c db_api_key;

DROP TABLE IF EXISTS api_key_users;
CREATE TABLE api_key_users(
    user_id VARCHAR(50) UNIQUE,
    api_key VARCHAR(150) UNIQUE,
    role VARCHAR(30)
);


