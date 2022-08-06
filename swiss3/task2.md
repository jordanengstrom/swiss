Now that we have our program cleanly refactored, let's add some new functionality. We'll start by allowing it to process more than a single resource. In process, remove the check that ensures that exactly one resource is supplied. Loop over all of the provided resources. Print out the name of the resource and then process it. If any of them fails, write out the error message, but continue processing any remaining files.

Change the Usage message for count to say "count the bytes for one or more resources," and change the Usage message for lang to say "find the language for one or more resources."

Build your program with go build -o swiss main.go.

Invoke this program two times with the parameters:

    count /tmp/file1.txt http://example.com /tmp/file2.txt /tmp/file3.txt
    lang /tmp/file1.txt http://example.com /tmp/file2.txt /tmp/file3.txt

Put the output for both runs into a file called swiss2.txt.

(Hint: you can use >> to append to an existing file.)
