#!/bin/bash
php generate.php
google-chrome --headless --disable-gpu --screenshot --window-size=320,420 output.svg