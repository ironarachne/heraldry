# Heraldry Generator

A tool to generate random heraldry.

Requires the following installed:

* PHP
* Composer
* php-gd
* php-xml

## Prepare

```
composer install
```

## Run

```
php generate.php
```

By default, this will create only a single pair of files (one for the image, and one for the blazon) in the `generated` directory. It will also output the XML and the blazon to STDOUT.

You can add a number to the end of the command to change the number of arms generated. For example, `php generate.php 3` will generate three sets of files, named consecutively.

If you're generating more than one image, you can also opt to add `silent` to the end of the command to get rid of the STDOUT output. For example: `php generate.php 7 silent`.

Alternatively, you can run `./getpng.sh` if you have Google Chrome installed to get a PNG named `screenshot.png`. This script takes a single argument which is the name of the image to use as input. So, for example:

```
./getpng.sh generated/arms-0.svg
```
