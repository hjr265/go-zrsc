package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	log.SetFlags(0)
	flag.Parse()
}

func main() {
	exePath, err := filepath.Abs(flag.Arg(0))
	catch(err)

	f, err := os.OpenFile(exePath, os.O_RDWR|os.O_APPEND, 0666)
	catch(err)

	zipw := zip.NewWriter(f)
	for _, name := range flag.Args()[1:] {
		err := filepath.Walk(name, func(name string, info os.FileInfo, err error) error {
			catch(err)

			zfh, err := zip.FileInfoHeader(info)
			catch(err)
			zfh.Name = name
			w, err := zipw.CreateHeader(zfh)
			catch(err)

			if !info.IsDir() {
				f, err := os.Open(name)
				catch(err)
				_, err = io.Copy(w, f)
				catch(err)
			}
			return nil
		})
		catch(err)
	}
	err = zipw.Close()
	catch(err)

	cmd := exec.Command("zip", "-A", flag.Arg(0))
	err = cmd.Run()
	catch(err)
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
