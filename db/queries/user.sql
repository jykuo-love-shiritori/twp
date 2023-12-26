-- name: UserGetInfo :one
SELECT "name",
    "email",
    "address",
    "image_id" as "image_url"
FROM "user" u
WHERE u."username" = $1;

-- name: UserUpdateInfo :one
UPDATE "user"
SET "name" = COALESCE($2, "name"),
    "email" = COALESCE($3, "email"),
    "address" = COALESCE($4, "address"),
    "image_id" = CASE
        WHEN sqlc.arg('image_id')::TEXT = '' THEN "image_id"
        ELSE sqlc.arg('image_id')::TEXT
    END
WHERE "username" = $1
RETURNING "name",
    "email",
    "address",
    "image_id" as "image_url";

-- name: UserGetPassword :one
SELECT "password"
FROM "user"
WHERE "username" = $1;

-- name: UserUpdatePassword :one
UPDATE "user"
SET "password" = sqlc.arg(new_password)
WHERE "username" = $1
RETURNING "name",
    "email",
    "address",
    "image_id" as "image_url";

-- name: UserGetCreditCard :one
SELECT "credit_card"
FROM "user"
WHERE "username" = $1;

-- name: UserUpdateCreditCard :one
UPDATE "user"
SET "credit_card" = $2
WHERE "username" = $1
RETURNING "credit_card";

-- name: AddUser :exec
INSERT INTO "user" (
        "username",
        "password",
        "name",
        "email",
        "address",
        "role",
        "credit_card",
        "enabled"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        'customer',
        '{}',
        TRUE
    );

-- name: UserExists :one
SELECT EXISTS (
        SELECT 1
        FROM "user"
        WHERE "username" = $1
            OR "email" = $2
    );

-- user can enter both username and email to verify
-- but writing "usernameOrEmail" is too long
-- name: FindUserInfoAndPassword :one
SELECT "username",
    "role",
    "password"
FROM "user"
WHERE "username" = $1
    OR "email" = $1;

-- name: SetRefreshToken :exec
UPDATE "user"
SET "refresh_token" = @refresh_token,
    "refresh_token_expire_date" = @expire_date
WHERE "username" = @username;

-- name: FindUserByRefreshToken :one
SELECT "username",
    "role"
FROM "user"
WHERE "refresh_token" = @refresh_token
    AND "refresh_token_expire_date" > NOW();

-- name: DeleteRefreshToken :exec
UPDATE "user"
SET "refresh_token" = NULL
WHERE "refresh_token" = @refresh_token;

-- name: AddShop :exec
INSERT INTO "shop" (
        "seller_name",
        "image_id",
        "name",
        "description",
        "enabled"
    )
VALUES(
        @seller_name,
        NULL,
        @name,
        '',
        FALSE
    );
