# Certificats
Certificats converter PEM to DER.

## Compiling

### Setup
- Install [Go](https://go.dev/doc/install) and [git](https://git-scm.com/downloads), then clone this project.
```bash
$ git clone https://github.com/InnoverseTeam/Certificats
$ cd Certificats
```

### Compiling using Go
To compile using go, you need to initialize the module, and you can run the project.
```bash
$ go mod init pemtoder
$ go run .
```

### Compiling using Docker
To compile using docker, you need to initialize the module, and you can run the project.
```bash
$ go mod init pemtoder
$ docker build -t pemtoder .
$ docker run -it --rm -v "C:\Users\your-username\Desktop\Certificats:/app" pemtoder /bin/bash
$ ./pemtoder /app/innoversecert.pem /app/innoversecert.der
$ ./pemtoder /app/innoversekey.pem /app/innoversekey.der
```
