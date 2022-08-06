package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/abadojack/whatlanggo"
	"github.com/example/swiss/count"
	"github.com/example/swiss/read"
	"github.com/urfave/cli/v2"
)

type Outputter interface {
	Add(resource string, output string)
	AddError(resource string, error string)
	String() string
}

type Result struct {
	Resource string `json:"resource"`
	Output   string `json:"output,omitempty"`
	Error    string `json:"error,omitempty"`
}

type StringOutput struct {
	// StringOutput will build up a slice of strings in the same format
	// as before (resource on one line, result or error on the next)
	Results []string `json:"results"`
}

func (so *StringOutput) Add(resource string, output string) {
	so.Results = append(so.Results, resource, output)
}

func (so *StringOutput) AddError(resource string, err string) {
	so.Results = append(so.Results, resource, err)
}

func (so *StringOutput) String() string {
	return strings.Join(so.Results[:], "\n")
}

type JSONOutput struct {
	// JSONOutput returns JSON output when String is called.
	Operation string   `json:"operation"`
	Results   []Result `json:"results"`
}

func (jo *JSONOutput) Add(resource string, output string) {
	jo.Results = append(jo.Results, Result{
		Resource: resource,
		Output:   output,
	})
}

func (jo *JSONOutput) AddError(resource string, err string) {
	jo.Results = append(jo.Results, Result{
		Resource: resource,
		Error:    err,
	})
}

func (jo *JSONOutput) String() string {
	out, err := json.Marshal(jo)
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

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

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "json",
			Usage: "write output as JSON.",
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
	outPutter := findOutputter(c)
	for _, resource := range c.Args().Slice() {
		resourceRc, err := toReadCloser(resource)
		if err != nil {
			outPutter.AddError(resource, err.Error())
			continue
		}
		result, err := command(resourceRc)
		if err != nil {
			outPutter.AddError(resource, err.Error())
			continue
		}
		outPutter.Add(resource, result)
	}
	fmt.Println(outPutter.String())
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

func findOutputter(c *cli.Context) Outputter {
	if c.IsSet("json") {
		return &JSONOutput{
			Operation: c.Command.Name,
		}
	}
	return new(StringOutput)
}
