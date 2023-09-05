#!/bin/bash

DIR=$1

for file in $(find $DIR -type f -iname "*.jpg")
do
    sips -s format heic "$file" --out "${file%.*}.heic" && exiftool -tagsFromFile "$file" -all:all "${file%.jpg}.heic" 
    rm "${file%.jpg}.heic_original"
    rm "$file"
done

for file in $(find $DIR -type f -iname "*.jpeg")
do
    sips -s format heic "$file" --out "${file%.*}.heic" && exiftool -tagsFromFile "$file" -all:all "${file%.jpeg}.heic" 
    rm "${file%.jpeg}.heic_original"
    rm "$file"
done

for file in $(find $DIR -type f -iname "*.bmp")
do
    sips -s format heic "$file" --out "${file%.*}.heic" && exiftool -tagsFromFile "$file" -all:all "${file%.bmp}.heic" 
    rm "${file%.bmp}.heic_original"
    rm "$file"
done

# Fix path is compiled in cmd in the parent folder
#fixPath $DIR
#fixMeta $DIR

