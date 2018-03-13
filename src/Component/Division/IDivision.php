<?php

namespace Heraldry\Component\Division;

use Heraldry\Component\Tincture;

interface IDivision {
    public function getBlazon();
    public function getElements( Tincture $tincture, $width, $height );
}