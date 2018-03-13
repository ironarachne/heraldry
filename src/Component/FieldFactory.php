<?php

namespace Heraldry\Component;

class FieldFactory {

    private $types = [
        'plain',
        'Pale',
        'Fess'
    ];

    public function generateRandom( $width, $height, $tincture1, $tincture2 = false ) {
        $type = $this->types[ mt_rand( 0, count( $this->types ) - 1 ) ];
        return new Field( $type, $width, $height, $tincture1, $tincture2 );
    }

}