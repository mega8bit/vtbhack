package main

import (
    "encoding/json"
    "errors"
    "fmt"
    tgbotapi "github.com/temamagic/telegram-bot-api"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
)

type SmsResponse struct {
    Status     string `json:"status"`
}

func notifyUser(userId uint64, msg string) bool {
    user, err := getUser(userId)
    if err != nil {
        return false
    }

    err = notifyUserBySms(user.Phone, msg)
    if err != nil {
        fmt.Println(err.Error())
    }
    err = notifyUserByTg(user.TelegramId, msg)
    if err != nil {
        fmt.Println(err.Error())
    }

    return true
}

func notifyUserBySms(phone int, msg string) error {
    requestUrl := "https://sms.ru/sms/send"

    Url, err := url.Parse(requestUrl)
    if err != nil {
        return err
    }

    parameters := url.Values{}
    parameters.Add("api_id", "3445df08-168e-2dc4-150c-e65f5f81fd1d")
    parameters.Add("to", strconv.Itoa(phone))
    parameters.Add("msg", msg)
    parameters.Add("json", "1")
    Url.RawQuery = parameters.Encode()

    requestUrl = Url.String()

    resp, err := http.Get(requestUrl)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    var res SmsResponse
    err = json.Unmarshal(body, &res)

    if res.Status == "OK" {
        return nil
    }
    return errors.New("status not OK")
}

func notifyUserByTg(tdUserId int, msg string) error {
    bot, err := tgbotapi.NewBotAPI("970879349:AAHUQ10UX9vjCDkA0g1seb6t1y3zFvjiCPw")
    if err != nil {
        return errors.New("status not OK")
    }
    _, err = bot.Send(tgbotapi.NewMessage(int64(tdUserId), msg))
    if err != nil {
        return errors.New("status not OK")
    }
    return nil
}
