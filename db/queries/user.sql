-- name: UserGetInfo :one

SELECT u.* FROM "user" u WHERE u.id = $1;

-- name: UserUpdateInfo :one

UPDATE "user"
SET
    "name" = COALESCE($2, "name"),
    "email" = COALESCE($3, "email"),
    "address" = COALESCE($4, "address"),
    "image_id" = COALESCE($5, "image_id")
WHERE "id" = $1
RETURNING *;

-- name: UserUpdatePassword :one

UPDATE "user"
SET
    "password" = sqlc.arg(new_password)
WHERE
    "id" = $1
    AND "password" = sqlc.arg(current_password)
RETURNING *;

-- name: UserGetCreditCard :one

SELECT "credit_card" FROM "user" WHERE "id" = $1 ;
