<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

class Fess implements IDivision {

    public function getBlazon() {
        return 'Per fess';
    }
    
    public function getElements( Tincture $tincture, $width, $height ) {
        $elements = [[
            'name' => 'rect',
            'attributes' => [
                'x' => '0',
                'y' => ceil( $height / 2 ),
                'width' => $width,
                'height' => ceil( $height / 2 ),
                'mask' => 'url(#shieldmask)',
                'style' => 'fill:' . $tincture->getCode(),
            ]
        ]];

        return $elements;
    }
}