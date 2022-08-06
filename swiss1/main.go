package main

import (
	"bytes"
	"errors"
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
			Usage:  "count the bytes in a resource",
			Action: counter,
		},
		{
			Name:   "lang",
			Usage:  "find the language of a resource",
			Action: langDetector,
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
	if c.Args().Len() != 1 {
		return errors.New("expected one resource")
	} else if c.Args().Len() == 1 {
		cmd, err := toReadCloser(c.Args().First())
		if err != nil {
			return err
		}
		// defer cmd.Close()

		res, err := command(cmd)
		if err != nil {
			return err
		}
		// defer res.Close()

		fmt.Println(res)
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
		// defer res.Close()

		return res, nil
	} else {
		res, err := read.FromFile(rp)
		if err != nil {
			return nil, err
		}
		// defer res.Close()

		return res, nil
	}
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
