## Single HLS Video Server
1. Converting original video file happens away from server.
2. Server serves .m3u8 files to clients via referenced .mp4.
3. Clients play HLS files using HLS.js or hls-video-element.


## Short term adds:
- Different qualities/resolution.
- Different captions for the file.
- Other basic video options to match.

## Long term TODOs:

UPLOAD 
- Add file uploader to update video directory if authenticated to do so.
- Add metadata attachments for certain fields (title, id)

ACCESS 
- Allow master access of mp4 file.
- Introduce signing to access a video.

TRANSCODING
- Implement concurrency for transcoding tasks.
- Implement caching for transcoded files.
- Implement error handling for transcoding tasks.
- Implement logging for transcoding tasks.
- Implement monitoring for transcoding tasks.
- Implement alerting for transcoding tasks.
- Implement rate limiting for transcoding tasks.
- Implement throttling for transcoding tasks.
- Implement retry mechanism for transcoding tasks.

## Command to convert video
```bash
bash scripts/convert.bash
```
