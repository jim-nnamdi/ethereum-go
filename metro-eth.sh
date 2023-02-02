#!/bin/bash

function ethLib(){
    echo $1
    make run-eth 
}
ethLib "running metro's ethereum-library!"
