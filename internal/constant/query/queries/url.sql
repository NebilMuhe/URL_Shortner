-- name: CreateURL :one
INSERT INTO urls (original_url, short_code) 
VALUES ($1, $2)
RETURNING *;

-- name: GetURLByShortCode :one
SELECT * FROM urls 
WHERE short_code = $1;

-- name: UpdateCount :one
UPDATE urls
SET count = count + 1
WHERE short_code = $1
RETURNING *;

-- name: UpdateURL :one
UPDATE urls
SET original_url = $1 and updated_at = NOW()
WHERE short_code = $2
RETURNING *;