<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

class BendSinister implements IDivision {

    public function getBlazon() {
        return 'Per bend sinister';
    }

    public function getElements( Tincture $tincture, $width, $height ) {
        $elements = [[
            'name' => 'polygon',
            'attributes' => [
                'points' => "0,0 $width,0 0,$height ",
                'class' => 'triangle',
                'mask' => 'url(#shieldmask)',
                'style' => 'fill:' . $tincture->getCode(),
            ]
        ]];

        return $elements;
    }

}