package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"io"
	"log"
	"time"
)

func constructBuildImageOptions(dockerfile []byte) (docker.BuildImageOptions, io.Reader) {
	t := time.Now()
	inputbuf, outputbuf := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	tr := tar.NewWriter(inputbuf)
	tr.WriteHeader(&tar.Header{Name: "Dockerfile", Size: 10, ModTime: t, AccessTime: t, ChangeTime: t})
	tr.Write(dockerfile)
	tr.Close()
	return docker.BuildImageOptions{
		Name:         "test",
		InputStream:  inputbuf,
		OutputStream: outputbuf,
	}, outputbuf

}

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	opts, outputbuf := constructBuildImageOptions([]byte("FROM base\n"))
	if err := client.BuildImage(opts); err != nil {
		log.Fatal(err)
	}
	fmt.Println(outputbuf)
}
