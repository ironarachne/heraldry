# Heraldry Generator

[![CircleCI](https://circleci.com/gh/ironarachne/heraldry.svg?style=svg)](https://circleci.com/gh/ironarachne/heraldry) [![Maintainability](https://api.codeclimate.com/v1/badges/a9824f8aa78dda6eb462/maintainability)](https://codeclimate.com/github/ironarachne/heraldry/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/a9824f8aa78dda6eb462/test_coverage)](https://codeclimate.com/github/ironarachne/heraldry/test_coverage)

A tool to generate random heraldry, written in Go.

## The markarms command

The binary output of this program is a command `makearms`. You can use this program to generate random heraldry in SVG format.

It has the following options:

`-html`: Generates an HTML index page displaying each image and its *blazon*.

`-results=N`: Generates `N` devices. The default is 1.

`-path=X`: Outputs all files to `X`. Defaults to the current directory.