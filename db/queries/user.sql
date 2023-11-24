-- name: UserGetInfo :one

SELECT u.* FROM "user" u WHERE u.id = $1;

-- name: UserUpdateInfo :one

UPDATE "user"
SET
    "name" = COALESCE($2, "name"),
    "email" = COALESCE($3, "email"),
    "address" = COALESCE($4, "address"),
    "image_id" = COALESCE($5, "image_id")
WHERE "id" = $1 RETURNING *;
