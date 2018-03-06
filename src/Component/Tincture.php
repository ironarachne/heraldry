<?php

namespace Heraldry\Component;

class Tincture {
    
    private $name;
    private $code;
    private $type;

    public function __construct( $name, $code, $type ) {
        $this->name = $name;
        $this->code = $code;
        $this->type = $type;
    }

    public function getName() {
        return $this->name;
    }

    public function getCode() {
        return $this->code;
    }

    public function getType() {
        return $this->type;
    }

}