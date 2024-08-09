-- name: CreateShop :exec
INSERT INTO shops (name, postal_code, city, street, building, phone_number, business_hours, description, latitude, longitude)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateShop :exec
UPDATE shops
SET name = ?, postal_code = ?, city = ?, street = ?, building = ?, phone_number = ?, business_hours = ?, description = ?, latitude = ?, longitude = ?
WHERE id = ?;

-- name: DeleteShop :exec
DELETE FROM shops WHERE id = ?;

-- name: CreateAdvertisementContent :exec
INSERT INTO advertisement_contents (shop_id, content, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- name: UpdateAdvertisementContent :exec
UPDATE advertisement_contents
SET content = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteAdvertisementContent :exec
DELETE FROM advertisement_contents WHERE id = ?;

-- name: CreateUser :exec
INSERT INTO users (name, email)
VALUES (?, ?);

-- name: UpdateUser :exec
UPDATE users
SET name = ?, email = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: CreateShopOwner :exec
INSERT INTO shop_owners (shop_id, name, email, password)
VALUES (?, ?, ?, ?);

-- name: GetShopOwnerById :one
SELECT * FROM shop_owners WHERE id = ?;

-- name: UpdateShopOwner :exec
UPDATE shop_owners
SET name = ?, email = ?, password = ?
WHERE id = ?;

-- name: DeleteShopOwner :exec
DELETE FROM shop_owners WHERE id = ?;

-- name: CreateReservation :exec
INSERT INTO reservations (shop_id, user_id, reservation_datetime, status)
VALUES (?, ?, ?, ?);

-- name: UpdateReservation :exec
UPDATE reservations
SET reservation_datetime = ?, status = ?
WHERE id = ?;

-- name: DeleteReservation :exec
DELETE FROM reservations WHERE id = ?;

-- name: AddFavoriteShop :exec
INSERT INTO favorite_shops (user_id, shop_id)
VALUES (?, ?);

-- name: RemoveFavoriteShop :exec
DELETE FROM favorite_shops WHERE user_id = ? AND shop_id = ?;

-- name: getFavoriteShops :many
SELECT * FROM favorite_shops AS f 
INNER JOIN shops as s 
ON f.shop_id = s.id 
WHERE f.user_id = ?
LIMIT 10 OFFSET ?;