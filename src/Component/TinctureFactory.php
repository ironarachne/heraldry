<?php

namespace Heraldry\Component;

class TinctureFactory {

    private $metals;
    private $colors;

    public function __construct() {
        $tinctureFileData = file_get_contents( 'metadata/tinctures.json' );
        $tinctureData = json_decode( $tinctureFileData );
        $this->metals = $tinctureData->metals;
        $this->colors = $tinctureData->colors;
    }
    
    private function generateRandom( $type ) {
        $typePlural = $type . 's';
        $randomElement = $this->$typePlural[ mt_rand( 0, count( $this->$typePlural ) - 1 ) ];

        $tincture = new Tincture( $randomElement->name, $randomElement->hexcode, $type );

        return $tincture;
    }

    public function randomColor() {
        return $this->generateRandom( 'color' );
    }

    public function randomMetal() {
        return $this->generateRandom( 'metal' );
    }
}