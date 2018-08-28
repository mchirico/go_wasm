<a href='https://jira.aipiggybot.io/projects/GWA/'>
<img src="https://storage.googleapis.com/montco-stats/JiraSoftware.png" alt="drawing" width="150px;"/>
         </a>


# go_wasm
Go Web Assembly



## Setup

I use the following program, in a parent directory, above this project. Note, the git clone command below -- this is above the project.

```code
#!/bin/bash

# Change Project Here:
PROJ=go_wasm

#
mkdir -p src/github.com
mkdir -p bin

export GOPATH=`pwd`
export PATH="$(pwd)/bin:$PATH"
export GOBIN="$(pwd)/bin"
if ! [ -x "$(command -v godep)" ]; then
    echo 'Note: godep is not installed.' >&2
    echo '... we will install it ..' >&2
    go get github.com/tools/godep

fi


if [ -d "$PWD/src/github.com/$PROJ" ]; then
    cd "$PWD/src/github.com/$PROJ"
else
    echo -e '

       cd src/github.com
       git clone git@github.com:mchirico/${PROJ}.git

'
fi


```


## Good Intro
[Add Function: Video](https://www.youtube.com/watch?v=4kBvvk2Bzis&feature=youtu.be)


