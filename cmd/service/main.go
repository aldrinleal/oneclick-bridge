package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aldrinleal/oneclick-bridge/types"
	"github.com/aldrinleal/oneclick-bridge/util"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	log "github.com/sirupsen/logrus"
)

var (
	sess *session.Session
	snsClient *sns.SNS
	topicArn string
)

func init() {
	sess = session.Must(session.NewSession(aws.NewConfig().WithRegion(util.EnvIf("AWS_DEFAULT_REGION", "AWS_REGION", "us-west-2"))))
	snsClient = sns.New(sess)
	topicArn = util.EnvIf("TOPIC_ARN", "")
}

func OneClickHandler(ctx context.Context, evt types.OneClickEvent) (err error) {
	buf := bytes.NewBuffer([]byte{})

	err = json.NewEncoder(buf).Encode(evt)

	if nil != err {
		log.Warnf("Oops: %s", err)

		return err
	}

	eventAsString := buf.String()

	log.Infof("eventsAsString: %s", eventAsString)

	_, err = snsClient.Publish(&sns.PublishInput{
		TopicArn: aws.String(topicArn),
		Message: aws.String(eventAsString),
	})

	if nil != err {
		log.Warnf("Oops: %s", err)
	}

	return err
}

func main() {
	lambda.Start(OneClickHandler)
}
