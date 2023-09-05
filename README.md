# Some Software for working with Photos

## Bash Scripts
These are scripts to convert files to a standard formate using the hardware encoders on an M1 mac.
```
convert_img.sh -- converts jpg,jpeg,bmp file formats to heic
convert_img_parallel.sh -- same as convert_img.sh but uses parallel to speed up opperation
convert_video.sh -- converts mp4,avi,mov,m2ts to mp4 hevc mp4 with acc audio
del.sh -- deletes various file formats that were found in scan
jpeg_convert.sh -- parallel convert for jpeg to heic. 
```

## Golang
### Consolidate
This recursively searches a folder, hashes all the files found, then produces a new 'consolidated' folder with all unique files.

## findExt
Find all extensions working with in a parent dir. Useful to convert to standard format

## findLost
This might be broken. But its designed to compare all unique files in one dir to another dir recursively. Useful to know if all files were imported correctly.

## fixMeta
Fix time stamp and other meta data. 

## fixPath
Convert file names to `hash.ext` for finding uniques.


