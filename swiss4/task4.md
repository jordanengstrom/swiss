For our final step, we are going to provide the option of writing our output in JSON format.

The structure of our JSON is as follows:

{
  "operation": "count",
  "results": [
    {
      "resource": "/tmp/file1.txt",
      "output": "1469"    
    },
    {
      "resource": "/tmp/file2.txt",
      "error": "open /tmp/file2.txt: no such file or directory"    
    }
  ]
}

The operation field specifies the command that was run. There is an entry in results for each resource supplied. The name of the resource is in the resource field. If the resource was processed successfully, the result is stored as a string in a field output. If there was an error during processing, the result is stored in a field called error.

UrFave CLI allows you to specify flags. In main, update the configuration to specify a Boolean flag with the name json and the usage "write output as JSON." For details on how to do this, look at the documentation for the Flags field on cli.App. If the flag is specified, the output will be in JSON. Otherwise, it will be text.

To implement the switch between text output and JSON output, define an interface called Outputter with three methods:

    Add(resource string, output string)
    AddError(resource string, error string)
    String() string

Define two structs that meet this interface. The first, StringOutput will build up a slice of strings in the same format as before (resource on one line, result or error on the next). When String is called, all of the strings will be returned.

The second struct, JSONOutput returns JSON output when String is called. Define two fields, Operation and Results. Operation is of type string and Results is of type []Result. In String, use Go's JSON support to turn the JSONOutput instance into a string.

Define a struct called Result to hold the Resource, Output, and Error data.

Write a function called findOutputter that takes in a*cli.Context and returns an Outputter. This function checks the *cli.Context to see if the json flag is set. If it is, it returns an instance of JSONOutput, initialized with the name of the command (which can be found in context.Command.Name). Otherwise, it returns a StringOutput.

Finally, modify processor. It calls findOutputter to get an Outputter instance. It will then invoke Add when a command runs successfully for a resource and AddError when a command returns an error for a resource. Once all resources are processed, processor prints the output of String.

Build your program with go build -o swiss main.go.

Invoke this program twice with the parameters:

    --json hash /tmp/file1.txt http://example.com /tmp/file2.txt /tmp/file3.txt
    count /tmp/file1.txt http://example.com /tmp/file2.txt /tmp/file3.txt

Put the output for both runs into a file called swiss4.txt.

(Hint: you can use >> to append to an existing file.)
