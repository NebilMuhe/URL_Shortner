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