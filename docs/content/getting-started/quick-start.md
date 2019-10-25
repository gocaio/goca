# Quick Start

GOCA&copy; works based in commands, some of them are used for proyect management like `database`, others gives information about GOCA&copy; configuration and finaly there are commandas to perform OSINT actions such as `dorker` and `scrapper`.

A simple dork for the term _Goca_:
```shell
>$ .goca dorker --term Goca
INFO[0000] [Staring the mighty Goca v0.3.0]
INFO[0000] [Selected engines: all]
INFO[0001] [Found a total of 14 links]
```

You can run a specific set of engines for analysing files.
```shell
>$ ./goca dorker --term Goca --engines Google,Bing
INFO[0000] [Staring the mighty Goca v0.3.0]
INFO[0000] [Selected engines: [Google Bing]]
INFO[0001] [Found a total of 14 links]
```

!!! note "Engines and Plugins"

    Seleccting engines will only run plugins build for that specifics engines.

On the other hand, you can scrap a domain to analyze all its public information by running:
```shell
>$ ./goca scrapper --domain goca.io
INFO[0000] [Staring the mighty Goca v0.3.0]
INFO[0000] [Scrapping  goca.io]
```
