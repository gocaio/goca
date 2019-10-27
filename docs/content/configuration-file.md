# Configuration file
Goca uses a simple comfiguration file, it is a `TOML` file. That format is quite similar to `INI` format.

We have decided to have an entry per each command that Goca is able to execute plus an entry fot the global options or flags. You can see an example below.

``` toml
[global]
useragent = "Goca Metadata Scanner"
loglevel = "info" # Values are: info (3), warn (2), error (1), debug (4)
threads = 1
plugins = ["all"] # Values are mimetype

[scrapper]
domain = "example.com"
depth = 1     # How many links levels to dig
threshold = 1 # Time between requests
save = false  # Save downloaded assets

[dorker]
term = "goca"
maxpages = 1      # Maximum number of pages
threshold = 1     # Time between requests
engines = ["all"] # Search engines to dork
save = false      # Save downloaded assets

[database]
file = ""

[plugin]
list = false
```

Goca will look up the configuration file in:

* `$HOME/<basefolder>/gocacfg.toml`
* `./gocacfg.toml`

## Working with the configuratoin file
The configuration file can be filled completely or just with those values that are static for you. You can run Goca without issue the predefined flags in the configuration file.

If you want to dork the term `Goca` you usually run Goca like this:

```shell
./goca dorker --term Goca
```

But, configure that term in the configuration file like this:

``` toml
[dorker]
term = "goca"
```

You will run Goca without `--term` flag instead.

```shell
./goca dorker
```
