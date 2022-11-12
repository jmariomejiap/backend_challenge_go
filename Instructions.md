# Instructions

## How to run the Web Server
To run the web server you have some options.

1. Run without building a binary file. To do so, on your terminal run the following command. ( You must be inside and at the root of the `backend_challenge_go` folder)
```
go run .
```

2. You can build and run the binary. The first part creates the binary, after that you can execute it. 
For example:
```bash
go build
./backend_challenge
```


## How to interact with the Web Server

The Web Server is listening on port `8080`. You can do a simple request  to the `/` endpoint and the web server will respond with some instructions and documentation.

To fully interact with the web server, you need to be able to send a csv file containing a matrix. For simplicity you make use of the `matrix.csv` file found on this project/folder.

### Steps

1. On one terminal make sure the web server is up and running. You should this message:
    ```bash
    "Im running on port 8080"
    ```

2. Interact with the web server.
    * Using Curl: (on another terminal)
        ```bash
        curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
        ```
    * Using Postman: Make sure you the body to `form-data`

    <img width="1060" alt="image" src="https://user-images.githubusercontent.com/22829270/201491047-1b05c6ab-b70e-428b-9ea2-d7153fa3ba10.png">


## Tests
This web server has a solid testing foundation. 
You can run its test by running these commands:
* short
    ```bash
    go test . 
    ```

* verbose
    ```bash
    go test . -v
    ```