package Service

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sharedSession *session.Session

func GetSession() (*session.Session, error) {
	if sharedSession == nil {
		var err error
		sharedSession, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region:      aws.String("us-east-1"),
				Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN")),
			},
		})
		if err != nil {
			return nil, err
		}
	}

	return sharedSession, nil
}

func UpdateSession(accessKeyId string, secretKey string, awsToken string) (errorGotten error) {
	var err error
	sharedSession, err = session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewStaticCredentials(accessKeyId, secretKey, awsToken),
		},
	})
	return err
}
