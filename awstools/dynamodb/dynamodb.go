package awstools

import (
	"strconv"

	awsSession "github.com/GuillemXanxo/go_tools/awstools/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	region             = "us-east-1"
	s3ConnectionString = "https://s3.amazonaws.com"
)

type DynamoDBClient struct {
	session  *session.Session
	dynamodb *dynamodb.DynamoDB
}

func NewDynamoDBClient() *DynamoDBClient {
	sess := awsSession.NewAWSSession(s3ConnectionString, region)
	dynamo := dynamodb.New(sess, &aws.Config{
		Region: aws.String(region),
	})

	return &DynamoDBClient{
		session:  sess,
		dynamodb: dynamo,
	}
}

//PERSON IS AN EXAMPLE OBJECT THAT IS USEFUL AS AN EXAMPLE
type Person struct {
	Id      int
	Name    string
	Website string
}

// PutItem inserts the struct Person
//Just tell the table you want to upload the object and use PutItem method
func (d *DynamoDBClient) PutItem(person Person) error {
	_, err := d.dynamodb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("TableName"),
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
			"Name": {
				S: aws.String(person.Name),
			},
			"Website": {
				S: aws.String(person.Website),
			},
		},
	})
	return err
}

// GetItem gets a Person based on the Id, returns a person
func (d *DynamoDBClient) GetItem(id int) (person Person, err error) {
	result, err := d.dynamodb.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{ // --> we can use Key because "id" is KeySchema for the table
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String("TableName"),
	})

	if err != nil {
		return person, err
	}
	err = dynamodbattribute.UnmarshalMap(result.Item, &person)
	return person, err
}

//Scan items from table --> We want to obtain ALL Persons called Marc through it key Name
func (d *DynamoDBClient) ScanItems() (listPersons []Person, err error) {
	input := &dynamodb.ScanInput{
		FilterExpression:          aws.String("contains(Name, :value)"), //This is the key we are contrasting
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":value": {S: aws.String("Marc")}},
		TableName:                 aws.String("TableName"),
	}
	output, err := d.dynamodb.Scan(input)
	if err != nil {
		return nil, err
	}
	listElements := output.Items
	err = dynamodbattribute.UnmarshalListOfMaps(listElements, listPersons)
	return listPersons, nil
}

// DeleteItem deletes a Person based on the Person.Id
func (d *DynamoDBClient) DeleteItem(personID int) error {
	_, err := d.dynamodb.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(personID)),
			},
		},
		TableName: aws.String("TableName"),
	})

	return err
}

// UpdateItem updates the Person based on the Person.Id
func (d *DynamoDBClient) UpdateItem(person Person) error {
	_, err := d.dynamodb.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("Name"),
			"#W": aws.String("Website"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Name": {
				S: aws.String(person.Name),
			},
			":Website": {
				S: aws.String(person.Website),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
		},
		TableName:        aws.String("TableName"),
		UpdateExpression: aws.String("SET #N = :Name, #W = :Website"),
	})

	return err
}
