We are going to base the work in our final challenge on what we did in the "Interacting with Third-Party Libraries Challenge". If you wrote this code before, you could reuse it, but it would be better to go back and refactor our application.

When you create this program with go mod init, call it github.com/example/swiss, since it's a like a Swiss Army Knife for working with local and remote files.

The suggested refactorings are as follows. Implementing them will make it easier to add features in future steps of this challenge:

    Write a function called toReadCloser that takes in a string and returns an io.ReadCloser and an error. If the passed-in string starts with http:// or https://, return the results of a call to read.FromWeb (from the "Interacting with Third-Party Libraries Challenge"). Otherwise, return the results of a called to read.FromFile.

    Update processor so that it has two parameters, a *cli.Context and a function parameter (which should look like command func(io.Reader) (string, error)) and returns an error. Check to make sure that exactly one resource is specified and return an error with the message "expected one resource" if this is not the case. Then call toReadCloser with that one resource. If an error is returned, return the error. If not, call command, passing it the io.ReadCloser. If an error is returned, return it. Otherwise, print the result and return nil.

    Write two functions, counter and langDetector. These functions have a single parameter of type *cli.Context and return an error. Call processor from these functions, passing in the *cli.Context and a function that takes in an io.Reader and returns a string and an error that implements the business logic for each function.
    In main, configure UrFave CLI to run as before. Rather than using processorWrapper as the function for the Action in both commands, use counter and langDetector.

Build your program with go build -o swiss main.go.

Invoke this program four times with the parameters:

    count http://example.com
    lang /tmp/file3.txt
    lang /tmp/file2.txt
    read /tmp/file3.txt

Put the output for all four runs into a file called swiss1.txt.

(Hint: you can use >> to append to an existing file.)
