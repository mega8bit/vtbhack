package main

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int
	Name         string
	Email        string
	Status       int
	Token        string
	PasswordHash string
	Phone        int
	TelegramId   int
	IsAdmin      bool
	Surname      string
	Photo        string
	Weight       int
	Position     string // Должность
}

func (User) TableName() string {
	return "user"
}

type Topic struct {
	Id            int
	Title         string
	TypeId        int
	StartDatetime time.Time
	EndDatetime   time.Time
	Status        int
}

func (Topic) TableName() string {
	return "topic"
}

type Alias = int
type TopicStatusStruct struct {
	Created     Alias
	InProcess   Alias
	NeedApprove Alias
	Closed      Alias
}

type Question struct {
	Id      int
	Title   string
	Status  int
	TopicId int
}

func (Question) TableName() string {
	return "question"
}

type Vote struct {
	Id         int
	QuestionId int
	UserId     int
	Result     int
}

func (Vote) TableName() string {
	return "vote"
}

type DbMessageWithUser struct {
	Id         int
	Body       string
	QuestionId int
	QuoteId    sql.NullInt64
	Name   string
}
type DbMessage struct {
	Id         int
	Body       string
	QuestionId int
	QuoteId    sql.NullInt64
	UserId     int
}


func (DbMessage) TableName() string {
	return "message"
}

type Document struct {
	Id      int
	File    string
	Name    string
	TopicId int
}

func (Document) TableName() string {
	return "document"
}
