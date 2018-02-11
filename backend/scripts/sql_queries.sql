CREATE DATABASE sponsorship_portal;

CREATE TABLE corporate_user (
    id SERIAL PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    portal_state JSON
);

CREATE TABLE participant (
    id SERIAL PRIMARY KEY,
    registration_id VARCHAR(255) NOT NULL,
    resume_text TEXT NOT NULL,
    document tsvector NOT NULL
);

INSERT INTO participant (id, registration_id, resume_text, document)
    VALUES (4, '54555gfhgf', 'resume_text', to_tsvector('resume_text' || '. ' || '[name]'));

CREATE INDEX idx_fts_search ON participant USING gin(document);

-- without ranking
SELECT registration_id, resume_text FROM participant WHERE document @@ to_tsquery('Travel | Cure');

-- with ranking
SELECT registration_id, resume_text, ts_rank(document, keywords) AS rank
    FROM participant, to_tsquery('java & python') keywords
    WHERE keywords @@ document
    ORDER BY rank DESC;