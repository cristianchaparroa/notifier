# Notifier

The Notifier is the service in charge to send notifications for statuses, news, etc.  

## Local Environment

The local environment is configured using docker-compose. For these purposes we should configure the .env file with the respective environment variables. 
Then we can run the service with docker-compose
```
docker-compose up
```

### Service testing.

1. If we want to send real emails we should configure all the SMTP env in the .env and set the APP_ENV=production.
2. In order to send a notification we should call the endpoint `/notifications`.

Parameters

| Parameter | Description                                   |   type |
|-----------|-----------------------------------------------|-------|
| content   | email content                                 | string |
| type      | status,news,marketing                         | string |
| recipient  | email where will be delivered the notification | string |

The following is a request example
```
curl -X POST http://localhost:8000/notifications  -d '{"content": "The challenge has been done. Would you mind reviewing it? ", "type": "news", "recipient": "daniela.valbuena@modak.live"}'
```

Responses

| code | Description                                                       | 
|------|-------------------------------------------------------------------|
| 201  | notification sent                                                 | 
| 400  | Bad request. The body request is not parseable                    | 
| 503  | The notification has been limited according to the business rules | 


## Design
The following diagram shows the system design

![title](images/notifier-schema.png)
