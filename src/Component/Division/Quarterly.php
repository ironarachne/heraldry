<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

class Quarterly implements IDivision {

    public function getElements( Tincture $tincture, $width, $height ) {
        $elements = [
            [
                'name' => 'rect',
                'attributes' => [
                    'x' => '0',
                    'y' => '0',
                    'width' => ceil( $width / 2 ),
                    'height' => ceil( $height / 2 ),
                    'mask' => 'url(#shieldmask)',
                    'style' => 'fill:' . $tincture->getCode(),
                ],
            ],
                [
                'name' => 'rect',
                'attributes' => [
                    'x' => ceil( $width / 2 ),
                    'y' => ceil( $height / 2 ),
                    'width' => ceil( $width / 2 ),
                    'height' => ceil( $height / 2 ),
                    'mask' => 'url(#shieldmask)',
                    'style' => 'fill:' . $tincture->getCode(),
                ],
            ],
        ];

        return $elements;
    }

}