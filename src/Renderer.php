<?php

namespace Heraldry;

use \Twig_Loader_Filesystem;
use \Twig_Environment;

class Renderer {

    private $loader;
    private $twig;
    
    public function __construct() {
        $this->loader = new Twig_Loader_Filesystem( 'templates' );
        $this->twig = new Twig_Environment( $this->loader );
    }

    public function renderIndex( $fileNames ) {
        
        foreach ( $fileNames as $fileName ) {
            $blazonFile = "generated/$fileName.blazon.txt";
            $blazon = file_get_contents( $blazonFile );

            $device = [
                'imageFileName' => $fileName,
                'blazon' => $blazon,
            ];

            $devices[] = $device;
        }

        $html = $this->twig->render( 'index.html.twig', [ 'devices' => $devices ] );

        $indexFile = fopen( 'generated/index.html', 'w' );

        fwrite( $indexFile, $html );
        fclose( $indexFile );
    }

    public function renderToFile( $heraldry, $fileName = 'output.svg' ) {
        $file = fopen( $fileName, 'w' );
        fwrite( $file, $heraldry->getSvg()->asXML() );
        fclose( $file );
        
        $blazonFile = fopen( "$fileName.blazon.txt", 'w' );
        fwrite( $blazonFile, $heraldry->getBlazon() );
        fclose( $blazonFile );
    }

    public function renderToConsole( $heraldry ) {
        echo $heraldry->getSvg()->asXML();
        echo $heraldry->getBlazon();
    }

}