<?php

require_once 'vendor/autoload.php';

use Heraldry\Heraldry;
use Heraldry\Renderer;

$imageWidth = 320;
$imageHeight = 420;
$iterations = 1;
$outputToConsole = true;
$fileNames = [];

$renderer = new Renderer();

if ( isset( $argv[ 1 ] ) ) {
    $iterations = $argv[ 1 ];
}

if ( isset( $argv[ 2 ] ) ) {
    if ( $argv[ 2 ] == 'silent' ) {
        $outputToConsole = false;
    }
}

for ( $i = 0; $i < $iterations; $i++ ) {
    $heraldry = new Heraldry( $imageWidth, $imageHeight );

    $fileName = "arms-$i.svg";

    $renderer->renderToFile( $heraldry, "generated/$fileName" );

    if ( $outputToConsole ) {
        $renderer->renderToConsole( $heraldry );
    }

    $fileNames[] = $fileName;
}

$renderer->renderIndex( $fileNames );