# Hello [Zrsc.go Example]

This example shows how to embed files in the binary and how to access them from within the program.

## Instructions

Build the program and embed the `.txt` files

    $ go build
    $ zrsc embed hello *.txt

Running the binary should now print the contents of `one.txt` and `two.txt`:
    
    $ ./hello
    Hello, lemonade!