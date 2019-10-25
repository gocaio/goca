# Install Goca

Goca is an open-source tool an can be acced to through it [GitHub](https://github.com/gocaio/goca) site. But if you don't care about source code and you only want to use Goca, just use one of the option that best fit you.

## From GitHub
Goca team generates a set of executables ready to be used. Those Goca binaries are avaliable from GitHub release page [https://github.com/gocaio/goca/releases](https://github.com/gocaio/goca/releases).

## From repository

## From source
This is the hard way to get Goca, usually this is for developers. However it does not matter if you are a developer or not, as far as you want to collaborate with us, obviosly we will appreciate a lot.

There are some requirements that needs to be meat before compiling Goca. First of all you need to have the Go compiler installed. You can follow instructions from the Golang official webpage [https://golang.org](https://golang.org).

Once Go is installed you have to clone the Goca reposity.
```shell
git clone git@github.com:gocaio/goca.git
```

Finally, you can build Goca just by running the following command from the project folder.
```shell
go build -o goca *.go
```

That command will produce a Goca binary called `goca` that you can just use.

If you can execute the produced binary you are ready to go. Use you favorite editor and issue your new features. The Goca team encourage you to share your features with the community.
