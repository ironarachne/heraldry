<?php

namespace Heraldry\Component;

class Division {
    private $elements;
    private $fieldWidth;
    private $fieldHeight;
    private $type;

    public function __construct( $type, $tincture, $width, $height ) {
        $this->type = $type;
        $this->fieldWidth = $width;
        $this->fieldHeight = $height;
        $this->tincture = $tincture;

        $this->generateElements();
    }

    public function generateElements() {
        if ( $this->type == 'pale' ) {
            $this->elements = [
                'name' => 'rect',
                'attributes' => [
                    'x' => ceil( $this->fieldWidth / 2 ),
                    'y' => '0',
                    'width' => ceil( $this->fieldWidth / 2 ),
                    'height' => $this->fieldHeight,
                    'mask' => 'url(#shieldmask)',
                    'style' => 'fill:' . $this->tincture->getCode(),
                ]
            ];
        } elseif ( $this->type == 'fess' ) {
            $this->elements = [
                'name' => 'rect',
                'attributes' => [
                    'x' => '0',
                    'y' => ceil( $this->fieldHeight / 2 ),
                    'width' => $this->fieldWidth,
                    'height' => ceil( $this->fieldHeight / 2 ),
                    'mask' => 'url(#shieldmask)',
                    'style' => 'fill:' . $this->tincture->getCode(),
                ]
            ];
        }
    }

    public function getElements() {
        return $this->elements;
    }
}