package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/kr/pretty"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("clement_AWS_REGION")),
		//Credentials: credentials.NewStaticCredentials(os.Getenv("clement_AWS_ACCESS_KEY_ID"), os.Getenv("clement_AWS_SECRET_ACCESS_KEY"), ""),
		MaxRetries: aws.Int(4),
	})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to create session")
		panic("fail to crate session")
	}
	svc := sqs.New(sess)

	// send the message
	sendParam := &sqs.SendMessageInput{
		MessageBody:  aws.String(uuid.New().String()),
		QueueUrl:     aws.String(os.Getenv("clement_AWS_QUEUE")),
		DelaySeconds: aws.Int64(3),
		// MessageGroupId:         aws.String(MessageGroup),
		// MessageDeduplicationId: aws.String(uuid.New().String()), // needs to be *string
	}
	sendResponse, err := svc.SendMessage(sendParam)
	if err != nil {
		fmt.Println(err)
		panic("failed to send message")
	}
	pretty.Println("sendming message: ", sendResponse)

	//receive thr message
	receiveParam := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(os.Getenv("clement_AWS_QUEUE")),
		MaxNumberOfMessages: aws.Int64(5),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(20),
		//AttributeNames: []*string{aws.String("Test"), // Required},
	}
	receiveResponse, err := svc.ReceiveMessage(receiveParam)
	if err != nil {
		fmt.Println(err)
		panic("failed to receive mesage")
	}
	pretty.Println("receiving message: ", receiveResponse)

	//delete the mesasge from sqs
	for _, message := range receiveResponse.Messages {
		deleteParam := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(os.Getenv("clement_AWS_QUEUE")),
			ReceiptHandle: message.ReceiptHandle,
		}
		_, err := svc.DeleteMessage(deleteParam)
		if err != nil {
			fmt.Println("failed to delete message from SQS queue")
			log.Println(err)
		}
		pretty.Println("Deleted message ID: ", *message.MessageId)
	}
}
