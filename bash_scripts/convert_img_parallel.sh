#!/bin/bash
set -Eeuo pipefail
DIR=$1

# Convert JPG images
convert_jpg_to_heic() {
    sips -s format heic "$1" --out "${1%.*}.heic" && exiftool -tagsFromFile "$1" -all:all "${1%.jpg}.heic" 
    rm "${1%.jpg}.heic_original"
    rm "$1"
}

JPG_FILES=$(find $DIR -type f -iname "*.jpg")
export -f convert_jpg_to_heic
echo "$JPG_FILES" | parallel --progress --eta convert_jpg_to_heic {} 


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



