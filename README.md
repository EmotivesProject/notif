# notif
## Introduction
Notif is the service that handles notifications, it's kind of split into two parts one part is in a form of a consumer and other services such as postit and chatter produce events for it to consume.

While the other part handles the requests from users to state that they've seen the notification.

## Production Environment Variables
```
DATABASE_URL - URL for the database. Should also be connecting to the uacl database
VERIFICATION_URL - This is the url of the uacl service, that way it can verify requests
HOST - In case the service needs to run on anything other than 0.0.0.0
PORT - In case the service needs to run on anything other than 80
NOTIFICATION_AUTH - Secret value that notif uses to verify requests
EMAIL_FROM - Email configuration.
EMAIL_PASSWORD - Email configuration.
EMAIL_LEVEL - What level of logs gets sent to the email address.
ALLOWED_ORIGINS - Cors setup.
```
## Endpoints
```
base URL is notif.emotives.net

GET - /healthz - Standard endpoint that just returns ok and a 200 status code. Can be used to test if the service is up
POST - /internal_notification - Used with notification_auth secret, it will create a notification from the request body
DELETE - /internal_notification/post/{id} - used with notification_auth secret, it will delete notifications relating to the post id. This will happen if a user deletes a post.
DELETE - /like/post/{id}/user/{username} - used with notification_auth secret, it will delete any like notification about the post id for the user. This is usually fired when someone unlikes a post to remove.

GET - /notification - creates an autologin token for a user. Usage of the token is unlimited and has no expire date.
POST - /notification/{id} - updates a notification id to seen.
POST - /notification/link/username/{username} - Updates all notifications for the user to seen that has the same link as the one in the request body. This is used when a user has multiple notifications about a post or user messages and visits just one.
```
## Database design
Uses a postgres database.
[See here for latest schema, uses the uacl_db](https://github.com/TomBowyerResearchProject/databases)