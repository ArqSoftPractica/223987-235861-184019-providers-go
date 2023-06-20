package Service

import (
	"223987-235861-184019-providers/Models"
	"223987-235861-184019-providers/Repository"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	CompanyQueueServiceActive = struct{ IsActive bool }{false}
)

type MessageData struct {
	Message string `json:"Message"`
}

func ReceiveCompanyMessages() {
	queueURL := os.Getenv("SQS_COMPANY_QUEUE_URL")
	params := &sqs.ReceiveMessageInput{
		AttributeNames:        aws.StringSlice([]string{"SentTimestamp"}),
		MaxNumberOfMessages:   aws.Int64(10),
		MessageAttributeNames: aws.StringSlice([]string{"All"}),
		QueueUrl:              &queueURL,
		VisibilityTimeout:     aws.Int64(30),
		WaitTimeSeconds:       aws.Int64(0),
	}
	for {
		sess, err := GetSession()
		if err != nil {
			log.Printf("Error creating AWS session: %s", err.Error())
			CompanyQueueServiceActive.IsActive = false
			time.Sleep(5 * time.Minute)
			continue
		}

		svc := sqs.New(sess)

		resp, err := svc.ReceiveMessage(params)
		if err != nil {
			log.Printf("Error receiving messages: %s", err.Error())
			CompanyQueueServiceActive.IsActive = false
			time.Sleep(5 * time.Minute)
			continue
		}

		CompanyQueueServiceActive.IsActive = true
		for _, msg := range resp.Messages {
			go func(msg *sqs.Message) {
				defer func() {
					_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
						QueueUrl:      &queueURL,
						ReceiptHandle: msg.ReceiptHandle,
					})
					if err != nil {
						log.Printf("Error deleting message: %s", err.Error())
						CompanyQueueServiceActive.IsActive = false
						time.Sleep(5 * time.Minute)
					}
				}()

				log.Println("Message Gotten", *msg.Body)

				var messageData MessageData
				if err := json.Unmarshal([]byte(*msg.Body), &messageData); err != nil {
					log.Printf("Error unmarshaling message data: %s", err.Error())
					CompanyQueueServiceActive.IsActive = false
					time.Sleep(5 * time.Minute)
					return
				}

				var companyToCreate Models.Company
				if err := json.Unmarshal([]byte(messageData.Message), &companyToCreate); err != nil {
					log.Printf("Error unmarshaling company message: %s", err.Error())
					CompanyQueueServiceActive.IsActive = false
					time.Sleep(5 * time.Minute)
					return
				}

				if err := upsertCompany(&companyToCreate); err != nil {
					log.Printf("Error upserting company: %s", err.Error())
					CompanyQueueServiceActive.IsActive = false
					time.Sleep(5 * time.Minute)
					return
				}
			}(msg)
		}
	}
}

func upsertCompany(company *Models.Company) error {
	Repository.UpsertCompany(company)
	return nil
}
