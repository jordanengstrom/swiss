package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/abadojack/whatlanggo"
	"github.com/example/swiss/count"
	"github.com/example/swiss/read"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.Commands = []*cli.Command{
		{
			Name:   "count",
			Usage:  "count the bytes for one or more resources",
			Action: counter,
		},
		{
			Name:   "lang",
			Usage:  "find the language for one or more resources",
			Action: langDetector,
		},
		{
			Name:   "hash",
			Usage:  "find the language for one or more resources",
			Action: hasher,
		},
	}

	app.ExitErrHandler = func(context *cli.Context, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	app.Run(os.Args)
}

func process(c *cli.Context, command func(io.Reader) (string, error)) error {
	for _, resource := range c.Args().Slice() {
		fmt.Println(resource)
		resourceRc, err := toReadCloser(resource)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := command(resourceRc)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
	return nil
}

func detect(r io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return "", err
	}

	rStr := buf.String()
	info := whatlanggo.Detect(rStr)

	lang := info.Lang.String()
	return lang, nil
}

func toReadCloser(rp string) (io.ReadCloser, error) {
	if rp[0:7] == "http://" || rp[0:8] == "https://" {
		res, err := read.FromWeb(rp)
		if err != nil {
			return nil, err
		}

		return res, nil
	} else {
		res, err := read.FromFile(rp)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func sha256hash(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	cksum := hex.EncodeToString(h.Sum(nil))
	return cksum, nil
}

func counter(c *cli.Context) error {
	err := process(c, count.FromReader)
	if err != nil {
		return err
	}
	return nil
}

func langDetector(c *cli.Context) error {
	err := process(c, detect)
	if err != nil {
		return err
	}
	return nil
}

func hasher(c *cli.Context) error {
	err := process(c, sha256hash)
	if err != nil {
		return err
	}
	return nil
}
