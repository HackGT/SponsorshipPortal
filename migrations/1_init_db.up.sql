BEGIN;
CREATE SCHEMA IF NOT EXISTS portal;

SET search_path TO portal, public;

CREATE TABLE IF NOT EXISTS sponsor_org (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sponsor_user (
    id SERIAL PRIMARY KEY,
    org_id INTEGER REFERENCES sponsor_org(id),
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS participant (
    id SERIAL PRIMARY KEY,
    registration_id TEXT NOT NULL,
    document tsvector
);

CREATE INDEX idx_participant_fts ON participant USING gin(document);

COMMIT;
