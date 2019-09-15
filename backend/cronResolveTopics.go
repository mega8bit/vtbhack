package main

import (
    "database/sql"
    "fmt"
    "time"
)

func cronResolveTopics() {
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
            t.status = 1
            AND t.end_datetime <= NOW()
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
                    CASE WHEN SUM(CASE WHEN za - protiv > 0 THEN 1 ELSE 0 END) > COUNT(qid)/2 THEN 1 ELSE 0 END AS topicRes
                FROM (
                    SELECT
                        q.id AS qid,
                        SUM(CASE WHEN result = 0 THEN 1 ELSE 0 END) AS za,
                        SUM(CASE WHEN result = 1 THEN 1 ELSE 0 END) AS vo,
                        SUM(CASE WHEN result = 2 THEN 1 ELSE 0 END) AS protiv
                    FROM topic t
                    INNER JOIN question q ON (t.id = q.topic_id)
                    INNER JOIN vote v ON (v.question_id = q.id)
                    WHERE t."id" = $1
                    GROUP BY qid
                ) qq
            `

            res := -1
            err := db.QueryRow(query, topic.Id).Scan(&res)
            if err == sql.ErrNoRows {
                continue
            }
            if err != nil {
                continue
            }

            status := 4
            if res == 1 {
                status = 3
            }

            _, err = db.Exec(`
                UPDATE topic SET 
                    status = $1
                WHERE id = $2
                `,
                status,
                topic.Id,
            )
        }

        time.Sleep(5 * time.Second)
    }
}