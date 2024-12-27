-- create user dev with password 'dev';
-- create database todolist;
-- grant all privileges on database todolist to dev;

DO $$
BEGIN
    -- Check if the database 'todolist' exists
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'todolist') THEN
        CREATE DATABASE todolist;
    END IF;

    -- Check if the user 'dev' exists
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'dev') THEN
        CREATE USER dev WITH PASSWORD 'dev';
    END IF;

    -- Grant privileges
    GRANT ALL PRIVILEGES ON DATABASE todolist TO dev;
END $$;