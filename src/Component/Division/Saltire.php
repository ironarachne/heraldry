<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

class Saltire implements IDivision {

    public function getBlazon() {
        return 'Per saltire';
    }

    public function getElements( Tincture $tincture, $width, $height ) {
        $midX = ceil( $width / 2 );
        $midY = ceil( $height / 2 );

        $elements = [
            [
                'name' => 'polygon',
                'attributes' => [
                    'points' => "0,0 $midX,$midY 0,$height",
                    'mask' => 'url(#shieldmask)',
                    'style' => 'fill:' . $tincture->getCode(),
                ],
            ],
                [
                'name' => 'polygon',
                'attributes' => [
                    'points' => "$width,0 $midX,$midY $width,$height",
                    'mask' => 'url(#shieldmask)',
                    'style' => 'fill:' . $tincture->getCode(),
                ],
            ],
        ];

        return $elements;
    }

}