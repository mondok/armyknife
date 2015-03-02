# armyknife
Random command line utilities for day to day work

## Setup
1.  Pull repository
2.  `go install`

## Commands
`army-knife w [URL]`
Performs a GET on a webpage and returns the contents of the URL.  Same thing as cURL basically.

`army-knife db`
Lists out the current Rails development database

`army-knife rmdd`
Remove Xcode derived data

`army-knife v [ENVIRONMENT_VARIABLE_NAME]`
Prints the value of an environment variable

`army-knife fs [ROOT_DIRECTORY] -p [PREPEND_TEXT] `
List all files and nested files starting at [ROOT_DIRECTORY]

`army-knife d [ROOT_DIRECTORY] -p [PREPEND_TEXT]`
List all directories and nested files starting at [ROOT_DIRECTORY] and prepend [PREPEND_TEXT]
