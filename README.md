# SKID-PDF

A microservice docker container for converting web pages into PDF files.

## Usage

### Render a PDF from a web page

``` bash
curl "http://localhost:8080/html?grayscale=false&landscape=true&uri=google.com" > google.pdf
```

Alternately, browse to ```http://localhost:8080/html?uri=https://github.com/vacoj/skid-pdf&portrait=0&grayscale=0``` with chrome to see the PDF.

### Using HTTP

#### GET: /html

**Streams back to the caller a PDF render of the website passed in the ```?uri={someurl}``` field.**

- ```?uri={google.com,http://google.com,https://google.com}``` can be any valid string representing an http endpoint.  http/https may be specified.
- ```?grayscale={true,false,1,0,T,F}``` determines whether or not the PDF will be created in grayscale.   Default is full color / grayscale false.
- ```?landscape={true,false,1,0,T,F}``` determines whether or not the PDF will be created in landscape mode.   Default is portrait mode / landscape false.

#### GET: /help

**301s you to the README.md on github.com**

#### GET: /

**301s you to the project repository on github.com**
<!--##### /gof

- ```?uri={google.com,http://google.com,https://google.com}``` can be any valid string representing an http endpoint.  http/https may be specified.  This method is untested and barely works.  Just left it in for anyone else who might want to play with it.-->

### Using AMQP

#### Submit a message for aysnchronous processing

``` javascript
{
    "url": "https://google.com",
    "grayscale": true,
    "landscape": true,
    "targetFileName": "test1.pdf",
    "targetFileDest": "./pdfs"
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

- open to suggestions