package main

import (
	"io"
	"os"

	"github.com/hjr265/go-zrsc/zrsc"
)

func main() {
	fOne, err := zrsc.Open("one.txt")
	catch(err)
	_, err = io.Copy(os.Stdout, fOne)
	catch(err)

	fTwo, err := zrsc.Open("two.txt")
	catch(err)
	_, err = io.Copy(os.Stdout, fTwo)
	catch(err)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
