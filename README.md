# armyknife
Random command line utilities for day to day work

## Setup
1.  Pull repository
2.  `go install`

## Commands
`army-knife web [URL]`
Performs a GET on a webpage and returns the contents of the URL.  Same thing as cURL basically.

`army-knife db`
Lists out the current Rails development database

`army-knife rmdd`
Remove Xcode derived data

`army-knife var [ENVIRONMENT_VARIABLE_NAME]`
Prints the value of an environment variable

`army-knife files [ROOT_DIRECTORY]`
List all files and nested files starting at [ROOT_DIRECTORY]

`army-knife files -p [PREPEND_TEXT] [ROOT_DIRECTORY]`
List all files and nested files starting at [ROOT_DIRECTORY] and prepend [PREPEND_TEXT****]


`army-knife dirs [ROOT_DIRECTORY]`
List all directories and nested files starting at [ROOT_DIRECTORY]
****
`army-knife dirs -p [PREPEND_TEXT] [ROOT_DIRECTORY]`
List all directories and nested files starting at [ROOT_DIRECTORY] and prepend [PREPEND_TEXT****]****