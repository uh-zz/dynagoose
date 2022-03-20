# What is?

This sample using github.com/guregu/dynamo.

We have a client to compare with the use of github.com/uh-zz/dynagoose.

# Prerequisites

### DynamoDB

Emulate dynamodb with docker for local verification

```
docker pull amazon/dynamodb-local
docker run -d --name dynamodb -p 8000:8000 amazon/dynamodb-local
```

### aws cli

Connect to local dynamodb with aws cli.

```
aws dynamodb list-tables --endpoint-url http://localhost:8000 --region ap-northeast-1
```

If the following message appears and fails

```
Unable to locate credentials. You can configure credentials by running "aws configure".
```

Specify Credential with `aws configure`
(For local verification, any value is acceptable)

```
aws configure
```

If OK, the following message appears

```
aws dynamodb list-tables --endpoint-url http://localhost:8000 --region ap-northeast-1
{
    "TableNames": []
}
```

### Create Table

Please Run this command

```
aws dynamodb create-table \
        --endpoint-url http://localhost:8000 \
        --table-name MyFirstTable \
        --attribute-definitions AttributeName=MyHashKey,AttributeType=S AttributeName=MyRangeKey,AttributeType=N \
        --key-schema AttributeName=MyHashKey,KeyType=HASH AttributeName=MyRangeKey,KeyType=RANGE \
        --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
```

After Run, Please check with the following command

```
aws dynamodb list-tables --endpoint-url http://localhost:8000 --region ap-northeast-1
{
    "TableNames": [
        "MyFirstTable"
    ]
}
```

# Usage

Run the command

```
DYNAMO_ENDPOINT=http://localhost:8000 go run client.go
```

Output:

```
read item: {MyHash 1 My First Text}
updated item: {MyHash 1 My Second Text}
deleted item: {MyHash 1 My Second Text}
```

# Testing

Run the docker

```
docker run -d --name dynamodb -p 8000:8000 amazon/dynamodb-local
```

If not create table yet

```
aws dynamodb create-table \
        --endpoint-url http://localhost:8000 \
        --table-name MyFirstTable \
        --attribute-definitions AttributeName=MyHashKey,AttributeType=S AttributeName=MyRangeKey,AttributeType=N \
        --key-schema AttributeName=MyHashKey,KeyType=HASH AttributeName=MyRangeKey,KeyType=RANGE \
        --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
```

```
go test
```
