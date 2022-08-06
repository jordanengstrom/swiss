Now that we can process multiple files, let's add an additional command to swiss. Calculating a hash of a file is a common operation, and Go includes hashing support in its standard library. We are going to add a new command called hash, which should follow the structure we have already implemented for our count and lang commands.

    Write a new function called hasher that uses the crypto/sha256 package to calculate a sha256 hash. Like counter and langDetector, the function should have a parameter of type *cli.Context and return an error. It should invoke processor with the *cli.Context and a function that takes in an io.Reader and returns a string and an error. This function implements the hashing logic (You can look at https://pkg.go.dev/crypto/sha256#example-New-File to get an idea of how to do this).

    Add a *cli.Command to main with the Name set to hash, the Usage set to "calculates the sha256 hash for one or more resources," and the Action set to hasher.

Build your program with go build -o swiss main.go.

Invoke this program with the parameters:

    hash /tmp/file1.txt http://example.com /tmp/file2.txt /tmp/file3.txt

Put the output into a file called swiss3.txt.
