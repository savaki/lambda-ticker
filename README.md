# lambda-ticker

lambda job designed to be triggered by a CloudWatch Scheduled Event on a 5m interval.

## Installation

This library uses the excellent apex package to manage the deployment of the lambda function. More
information on apex can be found at [http://apex.run/](http://apex.run/).

1. Create a role for the lambda function
2. Update project.json with the role arn; this lambda function role should have permissions to create the ticker topic
3. ```apex deploy```
4. From the AWS Console, add CloudWatch Scheduled Event event source that repeats on a 5m interval

## Overview

lambda-ticker is intended to create a stream of SNS events at a 5m interval.

``` 
CloudWatch -> Lambda -> SNS -> SQS
```

1. CloudWatch Schedule Event generates event on 5m interval
2. Event triggers lambda function
3. Lambda function upserts the ticker-5m topic and posts an event
4. SQS subscribers may listen to the event

## Event Format

``` json
{
  "event": "ticker:5m",
  "data": {
    "time": "2016-04-04T05:00:00Z"
  }
}
```
