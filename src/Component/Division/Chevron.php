<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

class Chevron implements IDivision {

    public function getElements( Tincture $tincture, $width, $height ) {
        $midX = ceil( $width / 2 );
        $midY = ceil( $height / 2 );

        $elements = [[
            'name' => 'polygon',
            'attributes' => [
                'points' => "0,$height $midX,$midY $width,$height",
                'class' => 'triangle',
                'mask' => 'url(#shieldmask)',
                'style' => 'fill:' . $tincture->getCode(),
            ]
        ]];

        return $elements;
    }

}