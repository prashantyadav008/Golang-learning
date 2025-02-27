<!-- @format -->

# Golang-learning

## Steps for Installing Go in Linux

### Remove Installed GoLang

#### Check Go Installed Place

    which go

#### Remove Path

    sudo rm -rf <path like - /usr/bin/go>

#### Remove Path

    $nano ~/.bashrc

##### Remove this Line

export PATH=$PATH:/usr/local/go/bin

and CTRL + X, phir Y, aur Enter.

    source ~/.bashrc

#### Check Go Version

    go version

### Install Latest Go Version

    wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz

    sudo tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz

    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

    source ~/.bashrc

#### Check Go Version and Path

    go version

    go env GOPATH

##### Also Checkout or Follow Go Documentaion

    https://go.dev/doc/install

## For Fresh Setup

    mkdir hello

    cd hello

    code .

#### Create Go Init File

    go mod init example.com/hello

#### Go File Create as “.go” extension

    touch hello.go

#### Create a Basic Function for Print "Hello World" and Run

      go run .

## Basic Commands in Go

#### Create Mod File

    go mod init example.com/<file_name>

#### Synchronize the example.com/<file_name> module's dependencies, adding those required by the code, but not yet tracked in the module.

    go mod tidy

#### Edit the example.com/<file_name> module to redirect Go tools from its module path

    go mod edit -replace example.com/greetings=../greetings
