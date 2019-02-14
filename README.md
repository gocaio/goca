![](img/Goca_banner.png)


**Goca** is a [FOCA](https://github.com/ElevenPaths/FOCA) fork written in Go, which is a tool used mainly to find metadata and hidden information in the documents its scans. These documents may be on web pages, and can be downloaded and analyzed with **Goca**.

It is capable of analyzing a wide variety of documents, with the most common being Microsoft Office, Open Office, or PDF files, although it also analyzes Adobe InDesign or SVG files, for instance.

These documents are searched for using search engines such as:

+ Google
+ Bing
+ DuckDuckGo
+ Yahoo
+ Ask

Then downloads the documents and extracts the EXIF information from graphic files, and a complete analysis of the information discovered through the URL is conducted even before downloading the file.

***

## Wiki

All the documentation, questions and instructions under the Wiki:

+ [Wiki](https://github.com/gocaio/Goca/wiki/0_Index)

***

## Contributing Guide

Please reade the Contributing guide:

+ [Contributing](CONTRIBUTING.md)

***

## USAGE

~~~bash
./Goca --domain hackercat.ninja --filetype=pdf,docx

  [+] Searching pdf in hackercat.ninja
    [i] Searching in Google...
    [i] Searching in Bing...
    [i] Searching in Yahoo...
    [i] Searching in DuckDuck Go...

  [+] Searching docx in hackercat.ninja
    [i] Searching in Google...
    [i] Searching in Bing...
    [i] Searching in Yahoo...
    [i] Searching in DuckDuck Go...

 [+] All documents downloaded to hackercat.ninja_Files

 [+] Analyzing metadata...

 [+] All files analyzed!
    [i] 32 pdf files
    [i] 13 docx files

 [+] Report saved as hackercat.ninja.json!
~~~

File types.
+ doc, docx
+ ppt, pptx
+ xls, xlsx
+ pdf
+ DS_Store
+ png, gif, jpg
+ .git/
+ mp3, mp4
+ All files: `--filetype-all`
