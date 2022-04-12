# go-rest-api-coding-challenge

## warning

This is the naive implementation, therefore this code should not be used in production environment. There is no warranty.

## motivation

There was a chance to try a coding challenge with `72 hours time limit.

## conclusion

Focus on clean code & clean design principles and best practices and work quality was the most important factor for me.
I decided to build a service from scratch with minimum third-party libraries. Just for fun.
Unfortunatelly, I had to pause my work after 48 hours. 
The result is tagged by `2_days_of_dev` git tag.

What remains:

 - [ ] https://github.com/lapacek/go-rest-api-coding-challenge/issues/1 implement SpaceX HTTP client request
 - [ ] https://github.com/lapacek/go-rest-api-coding-challenge/issues/2 implement a delete booking endpoint `(last remaining bonus point)`

Estimate: 5 hours remains to finish

## assignment

Imagine it’s 2049 and you are working for a company called SpaceTrouble that sends people to different places in our solar system. You are not the only one working in this industry. Your biggest competitor is a less known company called SpaceX. Unfortunately you both share the same launchpads and you cannot launch your rockets from the same place on the same day. There is a list of available launchpads and your spaceships go to places like: Mars, Moon, Pluto, Asteroid Belt, Europa, Titan, Ganymede. Every day you change the destination for all the launchpads. Basically on every day of the week from the same launchpad has to be a “flight” to a different place.

Information about available launchpads and upcoming SpaceX launches you can find by SpaceX API: https://github.com/r-spacex/SpaceX-API

Your task is to create an API that will let your consumers book tickets online.

In order to do that you have to create 2 endpoints:

Endpoint to book a ticket where client sends data like:
```
First Name
Last Name
Gender
Birthday
Launchpad ID
Destination ID
Launch Date
```

You have to verify if the requested trip is possible on the day from provided launchpad ID and do not overlap with SpaceX launches, if that’s the case then your flight is cancelled.

Endpoint to get all created Bookings.

Extra points:
When you use docker/docker-compose to run the project.
When you write unit/functional tests.
When you create an endpoint to delete booking.

Technical requirements:
Please, use Golang and Postgres.
Please, use github or bitbucket.
Commit your changes often. Do not push the whole project in one commit.

## build

```bash
$ docker-compose build
$ docker-compose up
```

## cleanup

Somtetimes you need.

```bash
$ docker-compose down --remove-orphans --volumes
```

## dockerization

There are two containers. One for a database and the second one for API. 

## database

The database is created automatically during the build. You can see the structure in the `db/` directory.

## testing

Business logic is covered by functional tests. There are also some unit tests that cover computations.
All tests are triggered automatically during the build process. If the test is not successful then the build process fails.

You can also run tests manually.

```bash
$ cd internal
$ go test -v ./...
```

I have started with manual testing of API endpoint with below command.

```bash
$ curl -v -X GET 0.0.0.0:8000/booking
```
