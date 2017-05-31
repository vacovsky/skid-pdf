# SKID-PDF

A microservice docker container for converting web pages into PDF files.

## Usage

### Render a PDF from a web page

``` bash
curl "http://localhost:8080/pdf?grayscale=false&landscape=true&uri=google.com" > google.pdf
```

Alternately, browse to ```http://localhost:8080/pdf?uri=https://github.com/vacoj/skid-pdf&portrait=0&grayscale=0``` with chrome to see the PDF.

### Using HTTP

#### GET: /pdf

Streams back to the caller a PDF render of the website passed in the ```?uri={someurl}``` field.

- ```?uri={google.com,http://google.com,https://google.com}``` can be any valid string representing an http endpoint.  http/https may be specified.
- ```?grayscale={true,false,1,0,T,F}``` determines whether or not the PDF will be created in grayscale.   Default is full color / grayscale false.
- ```?landscape={true,false,1,0,T,F}``` determines whether or not the PDF will be created in landscape mode.   Default is portrait mode / landscape false.

#### POST: /pdf

The POST method on this endpoint allows for more complicated query strings, and soon, headers and other form data to be sent to a target endpoint for PDF generation.

``` javascript
{
    "url": "https://requestb.in/17du8md1", // endpoint you want to turn into a PDF
    "data": "#safe=off&q=wkhtmltopdf", // should be in query string format: "?key1=somval&key2=anotherVal"
    "grayscale": true,
    "landscape": true,
    "headers": {
        "testheader1": "testheader1"
    },
    "postParams": {  // if any values are passed here, the request becomes a POST.  If you want a GET, use the "data" field to pass a query string.
        "testkey1": "testvalue1",
        "testkey2": "testvalue2"
    }
}
```

<!--// if action == "POST" or "PUT", this should be a json blob: "{\"key1\":\"someval\",\"key2\":\"anotherVal\"}"-->

#### GET: /help

301s you to the README.md on github.com

#### GET: /

301s you to the project repository on github.com - eventually this will become a little form to generate PDFs
<!--##### /gof

- ```?uri={google.com,http://google.com,https://google.com}``` can be any valid string representing an http endpoint.  http/https may be specified.  This method is untested and barely works.  Just left it in for anyone else who might want to play with it.-->

### Using AMQP

#### Submit a message for asynchronous processing

``` javascript
{
    "url": "https://requestb.in/17du8md1",  // endpoint you want to turn into a PDF
    "data": "#safe=off&q=wkhtmltopdf", // should be in query string format: "?key1=somval&key2=anotherVal"
    "grayscale": true,
    "landscape": true,
    "headers": {
        "testheader1": "testheader1"
    },
    "postParams": {  // if any values are passed here, the request becomes a POST.  If you want a GET, use the "data" field to pass a query string.
        "testkey1": "testvalue1",
        "testkey2": "testvalue2"
    },
    "targetFileName": "test1.pdf",  // name of file to be created
    "targetFileDest": "./pdfs" // destination folder for file to be placed.  Mounting a shared volume seems to be a way to export files for consumption elsewhere.  Might add more destination formats later.
}
```
The above message being posted on the queue will result in a  grayscale, landscape oriented PDF file being created at ```./pdfs/test1.pdf```, with content rendered from https://google.com.


## Setup

### Modify existing/create settings file

``` javascript
{
    "httpPort": "8080", // must be string; port for http synchronous work
    "useQueue": true, // if false, only HTTP synchronous access is turned on
    "queueConnectionString": "amqp://username:password@hostname:port/", // AMQP connection string.  Tested with RabbitMQ
    "queueChannel": "skidpdf", // name of queue to listen on
    "autoAck": true // Acknowledge receipt of enqueued messages
}
```

**By default, if no arguments are passed after the binary name, the default settings file provided is used.  To specify a settings file, run the binary/container like this:**

``` bash
/path/to/binary /path/to/settings.json
```

### Build the application
To compile the service and build the docker image, enter the following in your terminal starting in the root of this project (assuming you have docker set up and working):
``` bash 

# build the application
cd src/skid-pdf; go build;

# go back to root of the project
cd ../../
```

### Build the docker image for hosting

``` bash
# create docker image with all required elements in place
docker build -t skidpdf -f skid-pdf.dockerfile .

# turn the docker image on
docker run -p 8080:8080 -it -d skidpdf
```

## TODO:

- better support needed for controlling where and how asynchronously generated files are written
- find a way to sterilize inputs to prevent command-line injections. Alternately, wrap the wkhtmltopdf C library instead of the binary.


- open to other suggestions