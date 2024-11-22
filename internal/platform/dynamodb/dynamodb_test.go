package dynamodb_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDBConnection(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	svc := dynamodb.New(sess)

	input := &dynamodb.ListTablesInput{}

	result, err := svc.ListTables(input)
	if err != nil {
		t.Fatalf("failed to list tables: %v", err)
	}

	assert.NotNil(t, result)
	t.Logf("Tables: %v", result.TableNames)
}
