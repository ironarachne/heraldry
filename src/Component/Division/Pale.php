<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

class Pale implements IDivision {

    public function getElements( Tincture $tincture, $width, $height ) {
        $elements = [
            'name' => 'rect',
            'attributes' => [
                'x' => ceil( $width / 2 ),
                'y' => '0',
                'width' => ceil( $width / 2 ),
                'height' => $height,
                'mask' => 'url(#shieldmask)',
                'style' => 'fill:' . $tincture->getCode(),
            ]
        ];

        return $elements;
    }

}