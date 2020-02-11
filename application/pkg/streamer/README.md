# Streaming

## Types of sources
* Frame on request (New frames, are not got automaticly, each request initiate get process)
* Last valid frame (New frames automaticly got to cycle buffer, by request pointer returned)
* Video data subscription (New frames automaticly got to cycle buffer, reader subscribed to notification about new data arrival)
