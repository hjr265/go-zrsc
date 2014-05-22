# Server [Zrsc.go Example]

This example shows how to embed files in the binary and serve them over HTTP.

## Instructions

Build the program and embed the `.txt` files

    $ go build
    $ zrsc embed server public/one.txt public/two/three.txt

Running the binary should start a file server at `:8080`:
    
    $ ./server

You can then access the `.txt` files at [http://localhost:8080/public/one.txt](http://localhost:8080/public/one.txt) and [http://localhost:8080/public/two/three.txt](http://localhost:8080/public/two/three.txt).