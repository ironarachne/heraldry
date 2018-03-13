<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

class Bend implements IDivision {

    public function getElements( Tincture $tincture, $width, $height ) {
        $elements = [
            'name' => 'polygon',
            'attributes' => [
                'points' => "0,0 0,$height $width,$height",
                'class' => 'triangle',
                'mask' => 'url(#shieldmask)',
                'style' => 'fill:' . $tincture->getCode(),
            ]
        ];

        return $elements;
    }

}