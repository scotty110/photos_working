#!/bin/bash
#.m4v - Video
#.wmv - Video 
#.mp4 - Video 
#.avi - Video
#.mov - Video
#.m2ts - Video

DIR=$1

for file in $(find $DIR -type f -iname "*.mp4")
do
    cp $file /tmp/tmp_mp4_copy.mp4
    ffmpeg -i $file -c:v hevc_videotoolbox -q:v 63 -tag:v hvc1 -c:a aac /tmp/new_mp4.mp4
    mv /tmp/new_mp4.mp4 $file #"${file%.mp4}.mp4"
    exiftool -tagsFromFile /tmp/tmp_mp4_copy.mp4 $file 
    rm "${file%.mp4}.mp4_original"
    rm /tmp/*.mp4 
done

for file in $(find $DIR -type f -iname "*.avi")
do
    #ffmpeg -i $file -c:v libx265 -q:v 63 -tag:v hvc1 -c:a aac "${file%.avi}.mp4"
    ffmpeg -i $file -c:v hevc_videotoolbox -q:v 63 -tag:v hvc1 -c:a aac "${file%.avi}.mp4"
    exiftool -tagsFromFile $file "${file%.avi}.mp4"
    rm "${file%.avi}.mp4_original"
    rm $file
done

for file in $(find $DIR -type f -iname "*.mov")
do
    ffmpeg -i $file -c:v hevc_videotoolbox -q:v 63 -tag:v hvc1 -c:a aac "${file%.mov}.mp4"
    exiftool -tagsFromFile $file "${file%.mov}.mp4"
    rm "${file%.mov}.mp4_original"
    rm $file
done

for file in $(find $DIR -type f -iname "*.m2ts")
do
    #ffmpeg -i $file -c:v hevc_videotoolbox -q:v 63 -tag:v hvc1 -c:a copy "${file%.m2ts}.mp4"
    ffmpeg -i $file -c:v hevc_videotoolbox -q:v 63 -tag:v hvc1 -c:a aac "${file%.m2ts}.mp4"
    exiftool -tagsFromFile $file "${file%.m2ts}.mp4"
    rm "${file%.m2ts}.mp4_original"
    rm $file
done

# ffmpeg -loglevel panic