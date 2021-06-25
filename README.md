## Task

Build a rest API in the language of your choice with the following functionalities:
1. Save a new note
2. Update a previously saved note
3. Delete a saved note
4. Archive a note
5. Unarchive a previously archived note
6. List saved notes that aren't archived
7. List notes that are archived

The API will be used by a team of UX. 
There should be a full set of instructions on how to run it as well as
a motivation for the tech stack used and some proposed improvements.

## How to run:

## Solution:

### - API Design

### - Note model

### - Data storage
There are many ways this data can be persisted. I chose a NOSQL database for few reasons:
The data stored is very limited and not relational. Since our app has no authentication/authorization
and does not have multiple users or multiple relationships, it makes sense to store notes a JSON.

There are other ways to achieve the same results, e.g. a Redis db, DynamoDB (for the AWS lovers).
Ultimately I chose MongoDB because of the great support for its Go driver.

I created a cluster on my account with limited 512MB storage (free tier) which will be deleted after the review.
I could have also put the db inside a docker-compose command, 
but I think this solution has a lower risk of irritating the user.
Access is allowed from anywhere

### - Security and credentials


## Improvements