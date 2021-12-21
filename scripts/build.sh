#!/bin/bash

# Author: Chase Brand
# This script is distributed as a part of basebot. All rights reserved.

# CD to the directory of this script.
script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $script_dir

# Build basebot and move it to our root directory.
go build ../cmd/basebot
mv basebot ..

# Go back to our original directory.
cd $OLDPWD