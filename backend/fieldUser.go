package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Phone      int    `json:"phone"`
	TelegramId int    `json:"telegram_id"`
	IsAdmin    bool   `json:"is_admin"`
	Position   string `json:"position"`
	Status     int    `json:"status"`
	Weight     int    `json:"weight"`
	Photo      string `json:"photo"`
}

type UserList struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"surname": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.Int,
			},
			"telegramId": &graphql.Field{
				Type: graphql.Int,
			},
			"isAdmin": &graphql.Field{
				Type: graphql.Boolean,
			},
			"weight": &graphql.Field{
				Type: graphql.Int,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"photo": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var fieldUser = &graphql.Field{
	Type: userType,
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email, _ := p.Args["email"].(string)
		password, _ := p.Args["password"].(string)
		hasher := md5.New()
		hasher.Write([]byte(password))
		passwordHashStr := hex.EncodeToString(hasher.Sum(nil))
		user := &User{}
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
              email = $1
              AND password_hash = $2
        `, email, passwordHashStr).Scan(
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
            return nil, errors.New("Login or password is not valid")
        }
		if err != nil {
			return nil, err
		}
		return user, nil
	},
}

var fieldUserCreate = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"surname": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"phone": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"position": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"weight": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"telegramId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},

	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		var token = ""
		name, _ := p.Args["name"].(string)
		surname, _ := p.Args["surname"].(string)
		email, _ := p.Args["email"].(string)
		password, _ := p.Args["password"].(string)
		phone, _ := p.Args["phone"].(int)
		position, _ := p.Args["position"].(string)
		weight, _ := p.Args["weight"].(int)
		telegramId, _ := p.Args["telegramId"].(int)
		hasher := md5.New()
		hasher.Write([]byte(password))
		passwordHashStr := hex.EncodeToString(hasher.Sum(nil))
		hasher.Write([]byte(time.Now().Format("20060102150405")))
		token = hex.EncodeToString(hasher.Sum(nil))
		_, err := db.Exec(`
          INSERT INTO "user"(
            name, 
            surname, 
            email, 
            token,
            phone,
            weight,
            position,
            password_hash,
            telegram_id
            ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        `,
			name,
			surname,
			email,
			token,
			phone,
			weight,
			position,
			passwordHashStr,
			telegramId,
		)
		if err != nil {
			token = ""
		}
		return token, err
	},
}

var fieldUserEdit = &graphql.Field{
	Type: graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"surname": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"phone": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"position": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"weight": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"telegramId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},

	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)
		surname, _ := p.Args["surname"].(string)
		email, _ := p.Args["email"].(string)
		password, _ := p.Args["password"].(string)
		phone, _ := p.Args["phone"].(int)
		position, _ := p.Args["position"].(string)
		weight, _ := p.Args["weight"].(int)
		telegramId, _ := p.Args["telegramId"].(int)
		hasher := md5.New()
		hasher.Write([]byte(password))
		passwordHashStr := hex.EncodeToString(hasher.Sum(nil))
		hasher.Write([]byte(time.Now().Format("20060102150405")))
		_, err := db.Exec(`
          UPDATE "user" SET 
            name = $1, 
            surname = $2, 
            email = $3, 
            phone = $4,
            weight = $5,
            position = $6,
            password_hash = $7,
            telegram_id = $8
          WHERE id = $9
        `,
			name,
			surname,
			email,
			phone,
			weight,
			position,
			passwordHashStr,
			telegramId,
			id,
		)
		return true, err
	},
}

var userListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserList",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"surname": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var fieldUserAll = &graphql.Field{
	Type: graphql.NewList(userListType),

	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		rows, err := db.Query(`
          SELECT id, name, surname FROM "user" ORDER BY id desc
        `)
		if err != nil {
			return nil, err
		}
		var users []*UserList
		for rows.Next() {
			model := &UserList{}
			err = rows.Scan(&model.Id, &model.Name, &model.Surname)
			if err != nil {
				continue
			}
			users = append(users, model)
		}

		return users, err
	},
}