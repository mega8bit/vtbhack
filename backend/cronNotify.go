package main

import (
    "database/sql"
    "fmt"
    "time"
)

func cronNotify() {
    for {
        //0 - Создан
        //1 - Идёт голосование
        //2 - Требуется утверждение председателя
        //3 - Закрыт и принят
        //4 - Закрыт и не принят

        var topics []*Topic
        rows, err := db.Query(`
          SELECT 
            t.id,
            t.title,
            t.type_id,
            t.start_datetime,
            t.end_datetime,
            t.status,
            t.chairman_id
          FROM topic t
          WHERE
            t.status = 0
            AND t.start_datetime <= NOW()
        `)

        if err == sql.ErrNoRows {
            fmt.Println(err.Error())
            time.Sleep(5 * time.Second)
            continue
        }

        if err != nil {
            fmt.Println(err.Error())
            time.Sleep(5 * time.Second)
            continue
        }

        for rows.Next() {
            model := &Topic{}
            err = rows.Scan(
                &model.Id,
                &model.Title,
                &model.TypeId,
                &model.StartDateTime,
                &model.EndDateTime,
                &model.Status,
                &model.ChairmanId,
            )
            if err != nil {
                fmt.Println(err.Error())
                continue
            }
            topics = append(topics, model)
        }

        fmt.Println("topics cnt", len(topics))

        if len(topics) == 0 {
            time.Sleep(5 * time.Second)
            continue
        }

        for _, topic := range topics {
            query := `
                SELECT
                    u.id
                FROM "user" u
                INNER JOIN topic_user tu ON (tu.user_id = u."id")
                WHERE tu.topic_id = $1
            `

            rows, err = db.Query(query, topic.Id)
            if err == sql.ErrNoRows {
                continue
            }
            if err != nil {
                continue
            }

            for rows.Next() {
                var userId uint64
                err = rows.Scan(
                    &userId,
                )
                if err != nil {
                    fmt.Println(err.Error())
                    continue
                }
                if userId > 0 {
                    notifyUser(userId, "Началось голосование")
                }
            }

            _, err = db.Exec(`
                UPDATE topic SET 
                    status = $1
                WHERE id = $2
                `,
                1,
                topic.Id,
            )
        }

        time.Sleep(5 * time.Second)
    }
}