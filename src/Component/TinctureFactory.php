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
    
    private function generateRandom( $type, $skipTincture = false ) {
        $typePlural = $type . 's';

        $randomElement = $this->$typePlural[ mt_rand( 0, count( $this->$typePlural ) - 1 ) ];

        while ( $skipTincture && ( $skipTincture->getName() == $randomElement->name ) ) {
            $randomElement = $this->$typePlural[ mt_rand( 0, count( $this->$typePlural ) - 1 ) ];
        }

        $tincture = new Tincture( $randomElement->name, $randomElement->hexcode, $type );

        return $tincture;
    }

    public function random() {
        $types = [
            'color',
            'metal'
        ];
        $randomTinctureType = $types[ mt_rand( 0, 1 ) ];

        return $this->generateRandom( $randomTinctureType );
    }

    public function randomColor( $skipTincture = false ) {
        return $this->generateRandom( 'color', $skipTincture );
    }

    public function randomMetal( $skipTincture = false ) {
        return $this->generateRandom( 'metal', $skipTincture );
    }
}