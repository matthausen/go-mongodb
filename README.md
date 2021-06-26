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

Option 1:
- Install the dependencies with `go mod download`
- Use the command `go run cmd/main.go` from the repository root
- or build and run `go build -o main ./cmd/` and then `./build/main`

Option2:
- Run the docker container from my container registry `matthausen/thirdfort`
or build the image locally with the docker file provided and run it `docker build -t <image-name> .`

I have provided a Postman collection of endpoints and parameters to test the functionalities as JSON.
I used Postman, but you can use the client of your choice (e.g. Insomnia etc.)

To perform all actions on a note, simply pass a JSON that looks like this in the body of the request. 
```
{
    "text": "My special note"
    "archived": false
}
```
## Solution:

Even though there are 7 different requirements, in fact updating, archiving, unarchiving a note
are part of the same functionality which is updating a note. 
Therefore, there is really need for only 1 endpoint to handle all these 3 cases.

I decided to create 5 endpoints:
- create a note
- delete a note
- update a note (archive and unarchive too)
- list all notes
- list all archived notes

A note is represented by a struct with:
```
ID: primitive.ObjectID
Text: string
Archived: boolean
```

The app has been tested with Postman. Unit testing is unnecessary as there is no logic.
For further testing options please read the last paragraph. 

### - Why Go
A CRUD app can be built in virtually any language. I chose Go because of its readability, speed
and relatively small amount of boilerplate to get started. It also interacts perfectly with all types of databases 
without the need of heavy third party libraries, just small drivers.

### - Data storage
I chose a NOSQL database for few reasons:
The data stored is very limited and not relational. Since our app does not have multiple users or multiple relationships, it makes sense to store notes a JSON.

There are other ways to achieve the same results, e.g. Redis or DynamoDB (for the AWS lovers).
Ultimately I chose MongoDB because of the great support for its Go driver.

I created a cluster on my own account with limited 512MB storage (free tier) which will be deleted after the review.
I could have also put the db inside a docker-compose command, 
but I think this solution has a lower risk of irritating the user.
Access is allowed from anywhere.

### - Security and credentials
All sensitive data is stored inside an .env file. For this exercise the file has been included in the git history.
Obviously sensitive data should only be stored as environment variables in a CI/CD pipeline
or as kubernetes secrets. This is just for demonstration purposes.

## Improvements

Given more time I would have improved the app in the following ways:

- Merge the ListAll and ListAllArchived endpoint in 1 endpoint which will then filter out notes 
based on a parameter passed in the url by the user. This would further reduce code repetition.
  
- The endpoints are testable via Postman collection. Unit testing mongoDB is not necessary. 
  Integration testing, however, would be way more useful in this case, to guarantee that the actual database is always reachable. 
  An example of that would be to have a test database with mock entries.
  
