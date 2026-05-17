CREATE TABLE documents (
    id          SERIAL      PRIMARY KEY,
    name        TEXT        NOT NULL,
    data_format  TEXT       NOT NULL,
    data        JSONB       NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_documents_name ON documents (name);
CREATE INDEX idx_documents_data ON documents USING GIN (data);