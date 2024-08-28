-- name: UpdateUserBalance :exec
UPDATE users
SET balance = ?
WHERE id = ?;

-- name: RecordTransaction :exec
INSERT INTO transactions (sender_id, receiver_id, amount)
VALUES (?, ?, ?);

-- name: GetTransaction :one
SELECT * FROM transactions WHERE id = ? limit 1;