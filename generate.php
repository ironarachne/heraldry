<?php

require_once 'vendor/autoload.php';

use Heraldry\Heraldry;

$imageWidth = 320;
$imageHeight = 420;

$heraldry = new Heraldry( $imageWidth, $imageHeight );

$heraldry->renderToXML();
$heraldry->renderToFile();