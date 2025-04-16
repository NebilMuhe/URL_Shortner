CREATE TABLE urls (
    id UUID PRIMARY KEY gen_random_uuid(),
    original_url VARCHAR(255) UNIQUE NOT NULL,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    count int DEFAULT 0
    created_at timestampz DEFAULT NOW(),
    deleted_at timestampz NULL,
    updated_at timestampz NULL 
);
