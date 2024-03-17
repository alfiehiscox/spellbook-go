# Spellbook Go

A cached REST server for a wizard's spellbook.

As seens [here](https://www.youtube.com/watch?v=bFf-A27Rc9s&ab_channel=DreamsofCode).

## Endpoints

GET    /spells       =  Gets all spells
GET    /spells/<id>  =  Gets a single spell
POST   /spells       =  Creates a new spell
DELETE /spells/<id>  = Deletes a spell
PUT    /spells/<id>  = Updateda a spell

A Spell resource is a JSON object in the following form:

```json
{
    "id": 0,
    "name": "spell-name",
    "damage": 0,
    "created_at": "2024-00-00T00:00:00.000000Z",
    "updated_at": "2024-00-00T00:00:00.000000Z"
}
```

Created in 45 minutes, mostly googling

## Dependencies
- SQLite 3
- SQLC 
- Echo
