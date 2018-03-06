<?php

require_once 'vendor/autoload.php';

use Heraldry\Heraldry;

$imageWidth = 315;
$imageHeight = 425;

$heraldry = new Heraldry( $imageWidth, $imageHeight );
$heraldry->render();