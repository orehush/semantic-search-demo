CREATE TABLE IF NOT EXISTS synonyms_overriden (
    phrase TEXT PRIMARY KEY,
    synonyms JSONB NOT NULL
);
