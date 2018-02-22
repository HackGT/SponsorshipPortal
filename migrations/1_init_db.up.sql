BEGIN;

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS sponsor_orgs (
    id serial PRIMARY KEY,
    name text NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    deleted_at timestamptz
);

CREATE TRIGGER sponsor_orgs_modified_timestamp
    BEFORE UPDATE ON sponsor_orgs
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();

CREATE TABLE IF NOT EXISTS sponsors (
    id serial PRIMARY KEY,
    org_id integer REFERENCES sponsor_orgs(id),
    email text NOT NULL,
    password text NOT NULL,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    deleted_at timestamptz
);

CREATE TRIGGER sponsors_modified_timestamp
    BEFORE UPDATE ON sponsors
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();

CREATE TABLE IF NOT EXISTS participants (
    id serial PRIMARY KEY,
    registration_id text NOT NULL,
    document tsvector,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,
    deleted_at timestamptz
);

CREATE TRIGGER participants_modified_timestamp
    BEFORE UPDATE ON participants
    FOR EACH ROW
    EXECUTE PROCEDURE update_timestamp();

CREATE INDEX participant_full_text_index ON participants USING gin(document);

COMMIT;
