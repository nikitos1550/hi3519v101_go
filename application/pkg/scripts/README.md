# Scripting

Scripting is mostly planned now for setuping videopipeline and video processing, rather external interaction.

Typical cases for video processing:
* init videopipeline
* catch new raw frame, pass via several IVE processing steps and then pass to encoder
* put some overlay onto frame and send to encoder
* make some nnie calculations
* find QR code on the image
    * expose found qr data on image
    * expose via api?

## Execution model

* init script - one time run during startup
* onEvent script - run each time event occured
* onTimer - run each fixed period of time

* TODO how to store vars during several runs
* Some scripts can assume some backend (golang & c) structs, like nnie should be loaded into mem...

## Event types
* onNewRawData - run on new raw data availible 
* onNewEncodedData - run on new encoded data

Only one instance of each script can exist.
If script doesn`t finish and new same event occurs, another script run will not be fired.

## External world communications
How to expose result to external world?
* sendmail
* expose results as json via http api
* http request (libcurl)
* sample prints should be supported and should be accessible via http api

## Video processing concepts

* Data is not going through Lua, all processing is done in internals (go or c)
* Lua determines step of processing with ability to make decision branching based on previous steps
* ...

