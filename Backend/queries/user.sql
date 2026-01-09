-- name: CreateUser :exec
INSERT INTO users (
    username,
    passwdhash,
    email
) VALUES (
    $1, $2, $3
);

-- name: LoginUserWithUsername :one
SELECT id, username, passwdhash FROM users WHERE username = $1; 
