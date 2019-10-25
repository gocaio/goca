# Quick Start

Goca works based in commands, some of them are used for proyect management like `database`, others gives information about Goca configuration and finaly there are commandas to perform OSINT actions such as `dorker` and `scrapper`.

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

In both cases, scrapping and dorking, Goca accepts a plugin list flag so you can select what kind of metadata you want to extract and analyse.

The followinf command will run all plugins configured for `application/pdf` against PDF files.
```shell
>$ ./goca dorker --term Goca --plugins application/pdf
INFO[0000] [Staring the mighty Goca v0.3.0]
INFO[0000] [Selected all engines]
INFO[0001] [Found a total of 14 links]
```

Or, you might want to analyze JPG, DOC and PDF files.
```shell
>$ ./goca dorker --term Goca --plugins application/pdf,image/jpeg,application/msword
INFO[0000] [Staring the mighty Goca v0.3.0]
INFO[0000] [Selected all engines]
INFO[0001] [Found a total of 14 links]
```