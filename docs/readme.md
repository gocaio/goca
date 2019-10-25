# Goca Documentation Source Code

The documentation is build using [MKDocs](https://www.mkdocs.org/) with [material](https://squidfunk.github.io/mkdocs-material/) theme and some of its extensions.

To test and see the documentation locally just run the following command from this folder:
```
docker run --rm -it -p 8000:8000 -v ${PWD}:/docs squidfunk/mkdocs-material
```
