package main

import (
    "database/sql"
)

func getUser(userId uint64) (User, error) {
    user := User{}

    err := db.QueryRow(`
          SELECT 
            id, 
            name, 
            surname, 
            email, 
            status, 
            token,
            phone,
            telegram_id,
            is_admin,
            weight,
            position
        FROM "user"
            WHERE 
              id = $1
        `, userId).
    Scan(
        &user.Id,
        &user.Name,
        &user.Surname,
        &user.Email,
        &user.Status,
        &user.Token,
        &user.Phone,
        &user.TelegramId,
        &user.IsAdmin,
        &user.Weight,
        &user.Position,
    )
    if err == sql.ErrNoRows {
        return user, err
    }
    if err != nil {
        return user, err
    }

    return user, nil
}
