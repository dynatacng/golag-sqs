# golag-sqs

# Run
    go run main.go

Output :

## sending messages
sendming message:  &sqs.SendMessageOutput{
    _:                            struct {}{},
    MD5OfMessageAttributes:       (*string)(nil),
    MD5OfMessageBody:             &"fdfa9b3c8ee8fb861619bf116b9ffa5b",
    MD5OfMessageSystemAttributes: (*string)(nil),
    MessageId:                    &"54155f46-d052-428d-90ed-86c41fe4b651",
    SequenceNumber:               (*string)(nil),
}

## recevibing mnessages
receiving message:  &sqs.ReceiveMessageOutput{
    _:        struct {}{},
    Messages: {
        &sqs.Message{
            _:                      struct {}{},
            Attributes:             {},
            Body:                   &"here is the message body",
            MD5OfBody:              &"fdfa9b3c8ee8fb861619bf116b9ffa5b",
            MD5OfMessageAttributes: (*string)(nil),
            MessageAttributes:      {},
            MessageId:              &"54155f46-d052-428d-90ed-86c41fe4b651",
            ReceiptHandle:          &"AQEBwT6ADgsO7kWXLH7n+87v7xnjLwOobRDEdv/lJI4QxLmV+dmahkDie7E9WfK1IzPUCqmScxstyGv1w1TnvV5/zOZ4mglBosEGjQzmyc7sO5E9OEA+Q9R1eFRXHRi6Bs6uWQ8Q0qmMuz1+GoaQpEDbpAtz5fjKxD/HyP7JAaHF6GQZE9Q1RTjKYyHEi5YRRAkvGVzXU07jyyIgV8s5qo4bOHSOKHBs0iRG1eOiXobdk401NnbvBjEZ5dR79eew85e7rum5/nivmDio/sVPkMH7NuKVejmsZfXI0Ze8Gy2/o8HQc41H2XYrARPSMvXoq1+s5qCpECs3m2U8RMbjtRaM/OWe2fdB9omzQtx8madTpGrMSV+xT+4Gs7opz/4dipvkBlTl7BTngOogKohV7UArtg==",
        },
    },
}

## deleteing messages
Deleted message ID:  54155f46-d052-428d-90ed-86c41fe4b651

# 備註

> ReceiptHandle 是用來提供刪除 message 的參數而不是 message id，每一次取出來的值都會變，刪除時要提供最後取出來的 ReceiptHandle 值。


### 其他

1. 每一個 request 預設上限大小是 64K，最多可以到 256K
2. 一次收到的 message 預設是 1 筆，但你可以自已調整，最多 10 筆
3. 在 polling message 時，為了避免不停地取空的 message，建議使用 WaitTimeSeconds 參數(可以自已定義等待時間，最多 20 秒)
4. 當你取出來 message 後，這個 message 會暫時消失一段時間 (visibility timeout)，這段時間內不會再被取到，這個時間預設是 30 秒左右，你也可以自已設定。建議設定的時間要比你的 worker 處理這個 message 的時間長，避免在還沒處理完又再次取到。需要考量 network delay / slow machines / too much IO / 刪除 message 的時間 / etc. 再設定要給多長的時間
5. 當取出一個 message 且處理完後要記得 delete，如果沒有 delete，過了 visibility timeout 後，它還是會再被取出來