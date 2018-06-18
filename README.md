### TODO list sample web server

This application shows how we can build a simple todo list service using Go Lang.

#### Initial setup

First we need to install the app dependencies using [dep](https://golang.github.io/dep/docs/installation.html):
```
$ dep ensure
```

Then we need to define the environment variables to connect to our Postgres database in .bash_rc:
```
export TODO_DATABASE_HOST=127.0.0.1:5432
export TODO_DATABASE_USER=YOUR_USER
export TODO_DATABASE_PASSWORD=YOUR_PASS
export TODO_DATABASE_NAME=todo
```

#### Usage

No UI is available for this app so you need to use the Task service to manage your tasks:
```
# List of tasks
GET /task 
# Create task
POST /task
# Close task
PUT /task
# Cancel task
DELETE /task
```


