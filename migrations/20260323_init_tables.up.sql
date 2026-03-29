CREATE TABLE IF NOT EXISTS mlist.models (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    repository TEXT NOT NULL,
    version TEXT NOT NULL
);
