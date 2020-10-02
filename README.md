Code Challenge
- 
## Instruction
Lana has come to conclusion that users are very likely to buy awesome Lana merchandising from a physical store that sells the following 3 products: 

``` 
Code         | Name              |  Price
-----------------------------------------------
PEN          | Lana Pen          |   5.00€
TSHIRT       | Lana T-Shirt      |  20.00€
MUG          | Lana Coffee Mug   |   7.50€
```

Various departments have insisted on the following discounts:

 * The sales department thinks a buy 2 get 1 free promotion will work best (buy two of the same product, get one free), and would like this to only apply to `PEN` items.

 * The CFO insists that the best way to increase sales is with discounts on bulk purchases (buying x or more of a product, the price of that product is reduced), and requests that if you buy 3 or more `TSHIRT` items, the price per unit should be reduced by 25%.

Your task is to implement a simple checkout server and client that communicate over the network.

We'd expect the server to expose the following independent operations:

- Create a new checkout basket
- Add a product to a basket
- Get the total amount in a basket
- Remove the basket

The server must support concurrent invocations of those operations: any of them may be invoked at any time, while other operations are still being performed, even for the same basket.

At this stage, the service shouldn't use any external databases of any kind, but it should be possible to add one easily in the future.

Using Go, implement a checkout service and its client that fulfils these requirements.

Examples:

    Items: PEN, TSHIRT, MUG
    Total: 32.50€

    Items: PEN, TSHIRT, PEN
    Total: 25.00€

    Items: TSHIRT, TSHIRT, TSHIRT, PEN, TSHIRT
    Total: 65.00€

    Items: PEN, TSHIRT, PEN, PEN, MUG, TSHIRT, TSHIRT
    Total: 62.50€

**The solution should:**

- Be written in Go (let us know if this is your first time!)
- Build and execute in a Unix operating system.
- Focus on solving the business problem (less boilerplate!)
- Have a clear structure.
- Be easy to grow with new functionality.
- Not include binaries or dependencies.

**Bonus Points For:**

- Unit tests
- Functional tests
- Dealing with money as integers
- Formatting money output
- Useful comments
- Documentation
- Docker images / CI
- Commit messages (include .git in zip)
- Thread-safety
- Clear scalability


## Installation
You can use the solution in Docker or on your own Go environment:

## For Docker

For build the solution image
```
$ docker build -t go-lana . 
```
For run the solution image
```
$ docker run -d -p 8080:8080 go-lana 
```
For run API tests
```
$ docker run -it go-lana ./main.test
```

For run functional tests
```
$ docker run -it go-lana ./lana_challenge.test
```
And have fun!

## For Local Environment

Install dependencies
```
$ go mod download
```
Build app
```
$ go build cmd/checkout-api/main.go
```
Run Server
```
$ ./main
```
Run functional test
```
$ go test -v
```
Run API test
```
$ go test -c cmd/checkout-api/main_test.go cmd/checkout-api/main.go -v
```

## Consideration

- I tried to use SOLID Principles, I created a package for each "context"
- I have tried to create a solution that is an easy system to extend for new future promotions. In `internal/pricing/discount` we can add new promotions implementing promotions interface  
- Docker was included
- HTTP/REST with JSON support was added to allow the use of the service. You can use `requests.http` for test using Goland
- It's easy create new request `internal/handler/handler.go`

## API Documentation
The file `doc.html` has api documentation.

For generate api documentation:
```
$ docker run -it --rm -v $PWD:/doc quay.io/bukalapak/snowboard html -o doc.html input.apib
```


## Testing

If you use Goland/IntelliJ, you can use `requests.http` for testing the API



