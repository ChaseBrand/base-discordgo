@echo off

Rem Author: Chase Brand
Rem This script is distributed as a part of basebot. All rights reserved.

Rem Get the current directory for later.
set old_dir=%cd%

Rem CD to the directory of this script.
cd %~dp0

Rem Build basebot and move it to our root directory.
go build ../cmd/basebot
move basebot.exe %old_dir%

Rem Go back to our original directory.
cd %old_dir%