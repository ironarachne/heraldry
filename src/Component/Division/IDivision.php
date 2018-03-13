<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

interface IDivision {
    public function getElements( Tincture $tincture, $width, $height );
}