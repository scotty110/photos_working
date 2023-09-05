#!/bin/bash
set -Eeuo pipefail
DIR=$1

#Covert jpeg
convert_jpeg_to_heic() {
    sips -s format heic "$1" --out "${1%.*}.heic" && exiftool -tagsFromFile "$1" -all:all "${1%.jpeg}.heic" 
    rm "${1%.jpeg}.heic_original"
    rm "$1"
}

JPEG_FILES=$(find $DIR -type f -iname "*.jpeg")
export -f convert_jpeg_to_heic
echo "$JPEG_FILES" | parallel --progress --eta convert_jpeg_to_heic {}

#Convert BMP (not many of them)
for file in $(find $DIR -type f -iname "*.bmp")
do
    sips -s format heic "$file" --out "${file%.*}.heic" && exiftool -tagsFromFile "$file" -all:all "${file%.bmp}.heic" 
    rm "${file%.bmp}.heic_original"
    rm "$file"
done



