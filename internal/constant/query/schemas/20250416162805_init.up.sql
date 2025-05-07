CREATE TABLE urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    original_url VARCHAR(255) UNIQUE NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    count int DEFAULT 0,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    deleted_at timestamptz NULL,
    updated_at timestamptz NULL 
);
