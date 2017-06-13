# Cinema Network Api

### Default ENVs

* PORT=8080
* MONGODB_URL="mongodb://localhost:27017/cinema_network"

# Cinema Network API

### Import dependencies

* Install godep.  `go get github.com/tools/godep`
* From your project directory, Run `godep save`
* Build/Run your project as usual but add the -v to the go command
ifyou want to see it display the imports that it is using.
`go run -v main.go` (or) `go build -v`

### Build

Run `./build.sh`

or

`docker run --rm -v "$PWD":/go/src/github.com/dangelzm/cinema-network-api -w /go/src/github.com/dangelzm/cinema-network-api iron/go:dev go get github.com/tools/godep && godep go build -o myapp`

then

`docker build -t ivanzmerzlyi/cinema-network-api:latest .`

### Run

`docker run --rm -it -p 3030:8080 ivanzmerzlyi/cinema-network-api --name cinema-network-api`

### Run local with Docker Composer

`docker-compose up`
