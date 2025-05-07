-- name: CreateURL :one
INSERT INTO urls (original_url, short_code) 
VALUES ($1, $2)
RETURNING *;