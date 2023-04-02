# check if yt-dlp is installed
if [ -f alarm ] &&  command -v yt-dlp &> /dev/null; then
    yt-dlp `./alarm`
elif [ -f alarm ]; then
    echo "Please install yt-dlp before running this script"
elif [ ! -f alarm ]; then
    echo "please run \"go build\" to compile before running this script"
else 
    echo "I'm confused"
fi