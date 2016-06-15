<?php

//
// This is only a SKELETON file for the "Hello World" exercise.
// It's been provided as a convenience to get you started writing code faster.
//

function helloWorld($name = null)
{
    // var_dump($name);
    if (is_null($name)) {
        return 'Hello, World!';
    } else {
        return 'Hello, ' . $name . '!';
    }
}
