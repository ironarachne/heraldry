# Heraldry Generator

A tool to generate random heraldry, written in Go.

## The makearms command

The binary output of this program is a command `makearms`. You can use this program to generate random heraldry in SVG format.

It has the following options:

`-html`: Generates an HTML index page displaying each image and its *blazon*.

`-results=N`: Generates `N` devices. The default is 1.

`-path=X`: Outputs all files to `X`. Defaults to the current directory.
