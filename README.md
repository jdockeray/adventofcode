### Advent of code

Requirements
--
- Docker for desktop
- Docker compose 

Installation and Running 
--
To run, build the services

```docker-compose build```

then start the file server (contains the input from advent of code)

```docker-compose run files```

Now you can run the app service against a given day, by changing the DAY environment variable

```
docker-compose run -e DAY=day1 app    
```
