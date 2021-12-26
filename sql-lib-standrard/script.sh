#!/bin/bash

TARGET=$1

curl "http://localhost:4444/weather?city=$TARGET"
