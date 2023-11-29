-- name: AddUser :exec

INSERT INTO
    "user" (
        "username",
        "password",
        "name",
        "email",
        "address",
        "image_id",
        "role",
        "credit_card",
        "enabled"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        '',
        $5,
        'customer',
        '{}',
        TRUE
    );

-- name: UserExists :one

SELECT EXISTS (
        SELECT 1
        FROM "user"
        WHERE
            "username" = $1
            OR "email" = $2
    );

-- name: FindUserPassword :one

SELECT "password" FROM "user" WHERE "username" = $1 OR "email" = $1;
