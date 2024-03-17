-- name: GetSpells :many 
SELECT * FROM spells; 

-- name: GetSpell :one
SELECT * FROM spells 
WHERE id = ? LIMIT 1;

-- name: DeleteSpell :exec 
DELETE FROM spells 
WHERE id = ?; 

-- name: CreateSpell :one
INSERT INTO spells (name, damage) 
VALUES (?, ?) 
RETURNING *; 

-- name: UpdateSpell :one
UPDATE spells 
SET name = ?, damage = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;
