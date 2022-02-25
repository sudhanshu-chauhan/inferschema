### Infer Schema Application
A sample rest api application exposing an endpoint to determine schema of an api

### Dependencies
- Linux OS (debian)
- Docker
- Supervisor
- go1.17.6

### Golang Dependencies
- github.com/gorilla/mux

### Running Tests
```
go test inferschema/app
```

### Installation Steps
 ```shell
 docker build -t <app>:<tag> .
 docker run -d -p 8000:8000 <app>:<tag>
 ```


 Assertions & Caveats
 
 - Only first row is parsed for schema determination, all rows can be parsed but will require more time for same.
 - The Datetime type parsing can be unreliable for a case like `02/02/2020` where we are not sure if it is `dd/mm/yyyy` or `mm/dd/yy`.
 - The Autmoatic restart in event of crash is being tackled by `Supervisor` daemon.
 - The deployment can be handled by registering the docker file into AWS ECR and from there, pushing into an AWS ECS.