<?php

namespace Heraldry;

use SVG\SVGImage;
use SVG\Nodes\Shapes\SVGRect;
use Heraldry\Component\TinctureFactory;

class Heraldry {

    private $image;
    private $height;
    private $width;

    public function __construct( $width, $height ) {
        $this->width = $width;
        $this->height = $height;

        $this->image = new SVGImage( $this->width, $this->height );

        $document = $this->image->getDocument();

        $background = new SVGRect( 0, 0, $this->width, $this->height );
        $background->setStyle( 'fill', '#FFFFFF' );
        $document->addChild( $background );

        $this->addShield();
    }

    public function addShield() {
        $shieldSource = SVGImage::fromFile( './source-images/shield.svg' );

        $shield = $shieldSource->getDocument()->getChild( 0 );

        $tinctureFactory = new TinctureFactory();
        $tincture = $tinctureFactory->random();

        $shield->setStyle( 'fill', $tincture->getCode() );

        $document = $this->image->getDocument();

        $document->addChild( $shield );
    }

    public function render() {
        $rasterImage = $this->image->toRasterImage( $this->width, $this->height );

        imagepng( $rasterImage, 'shield.png' );
    }

}