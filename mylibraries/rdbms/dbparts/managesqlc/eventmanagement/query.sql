-- name: CreateEvent :exec
INSERT INTO events (name, description, start_time, end_time, location) 
VALUES ($1, $2, $3, $4, $5);

-- name: GetEventById :one
SELECT * FROM events WHERE id = $1;

-- name: UpdateEvent :exec
UPDATE events 
SET name = $2, description = $3, start_time = $4, end_time = $5, location = $6 
WHERE id = $1;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;
