package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/abadojack/whatlanggo"
	"github.com/example/processor/count"
	"github.com/example/processor/read"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.Commands = []*cli.Command{
		{
			Name:   "count",
			Usage:  "count the bytes in a resource",
			Action: processWrapper,
		},
		{
			Name:   "lang",
			Usage:  "find the language of a resource",
			Action: processWrapper,
		},
	}

	app.ExitErrHandler = func(context *cli.Context, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	app.Run(os.Args)
}

func process(cmd, rp string) (string, error) {
	if rp[0:7] == "http://" || rp[0:8] == "https://" {
		res, err := read.FromWeb(rp)
		if err != nil {
			return "", err
		}
		defer res.Close()

		if cmd == "count" {
			n, err := count.FromReader(res)
			if err != nil {
				return "", err
			}
			return strconv.Itoa(n), nil
		} else if cmd == "lang" {
			l, err := detect(res)

			if err != nil {
				return "", err
			}

			return l, nil
		} else {
			return "", errors.New("unknown command")
		}

	} else {
		res, err := read.FromFile(rp)
		if err != nil {
			return "", err
		}
		defer res.Close()

		if cmd == "count" {
			n, err := count.FromReader(res)
			if err != nil {
				return "", err
			}

			return strconv.Itoa(n), nil
		} else if cmd == "lang" {
			l, err := detect(res)

			if err != nil {
				return "", err
			}

			return l, nil
		} else {
			return "", errors.New("unknown command")
		}
	}
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

func processWrapper(c *cli.Context) error {
	if c.Args().Len() > 1 {
		return cli.Exit("expected one resource", 1)
	} else if c.Args().Len() == 1 {
		str, err := process(c.Command.Name, c.Args().First())
		if err != nil {
			return cli.Exit(err, 1)
		}
		fmt.Println(str)
	}
	return nil
}
