<?php

namespace Heraldry;

use Heraldry\Component\FieldFactory;
use Heraldry\Component\TinctureFactory;

class Heraldry {

    private $field;
    private $height;
    private $svg;
    private $tincture1;
    private $tincture2;
    private $tincture3;
    private $width;

    public function __construct( $width, $height ) {
        $this->width = $width;
        $this->height = $height;

        $this->svg = simplexml_load_file( './source-images/shield.svg' );

        $this->generateTinctures();
        $this->addField();
        $this->renderElements();
    }

    private function addField() {
        $fieldFactory = new FieldFactory();
        $this->field = $fieldFactory->generateRandom( $this->width, $this->height, $this->tincture1, $this->tincture2 );
    }

    private function generateTinctures() {
        $tinctureFactory = new TinctureFactory();
        $this->tincture1 = $tinctureFactory->randomColor();
        $this->tincture2 = $tinctureFactory->randomColor();
        $this->tincture3 = $tinctureFactory->randomMetal();
    }

    public function getBlazon() {
        $blazon = $this->field->getBlazon();
        $blazon = ucfirst( $blazon );
        return $blazon;
    }

    public function getSvg() {
        return $this->svg;
    }

    public function renderElements() {
        $elements = $this->field->getElements();

        foreach ( $elements as $element ) {
            $child = $this->svg->addChild( $element[ 'name' ] );
            foreach ( $element[ 'attributes' ] as $attribute => $value ) {
                $child->addAttribute( $attribute, $value );
            }
        }
    }

}