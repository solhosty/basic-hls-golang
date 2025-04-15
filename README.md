## Single HLS Video Server
1. Encoding/Transcoding happens away from server.
2. Server serves HLS files to clients.
3. Clients play HLS files using HLS.js or hls-video-element.

## Short term adds:
- Different qualities/resolution.
- Different captions for the file.
- Other options match.

## Long term TODOs:
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
