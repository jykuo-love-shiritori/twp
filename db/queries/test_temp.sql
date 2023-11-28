-- name: InsertTestUser :exec

INSERT INTO
    "user" (
        "username",
        "password",
        "name",
        "email",
        "address",
        "image_id",
        "role",
        "credit_card"
    )
VALUES (
        'user0',
        'password0',
        'name0',
        'email0',
        'address0',
        $1,
        'customer',
<<<<<<< HEAD
        '{"card_number": "**** **** **** 7890",
        "expiry_date": "01/23",
        "name": "card 5",
        "CVV": "053"}'
=======
        '{"card_number": "card_number0", "expiration_date": "expiration_date0", "cvv": "cvv0"}'
>>>>>>> twp/main
    );

-- name: DeleteTestUser :exec

DELETE FROM "user" WHERE "username" = ' user0 ';

-- name: SearchTestUser :one

SELECT * FROM "user" WHERE "username" = ' user0 ';
