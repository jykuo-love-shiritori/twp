-- name: UserGetInfo :one

SELECT
    "id",
    "name",
    "email",
    "image_id",
    "enabled"
FROM "user" u
WHERE u.id = $1;

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
RETURNING
    "id",
    "name",
    "email",
    "address",
    "image_id",
    "enabled";

-- name: UserGetCreditCard :one

SELECT "credit_card" FROM "user" WHERE "id" = $1;

-- name: UserUpdateCreditCard :one

UPDATE "user"
SET "credit_card" = $2
WHERE "id" = $1
RETURNING
    "id",
    "name",
    "email",
    "address",
    "image_id",
    "enabled";
