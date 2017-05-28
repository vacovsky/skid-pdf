# SKID-PDF

A microservice docker container for converting web pages into PDF files.

## Basic Usage

### Render a PDF from a web page

``` bash
curl "http://localhost:8080/html?grayscale=false&landscape=true&uri=google.com" > google.pdf
```

Alternately, browse to ```http://localhost:8080/html?grayscale=false&landscape=true&uri=google.com``` with chrome to see the PDF.

### Endpoints

#### /html

- ```?uri={google.com,http://google.com,https://google.com}``` can be any valid string representing an http endpoint.  http/https may be specified.
- ```?grayscale={true,false,1,0,T,F}``` determines whether or not the PDF will be created in grayscale.   Default is full color / grayscale false.
- ```?landscape={true,false,1,0,T,F}``` determines whether or not the PDF will be created in landscape mode.   Default is portrait mode / landscape false.

##### /gof

- ```?uri={google.com,http://google.com,https://google.com}``` can be any valid string representing an http endpoint.  http/https may be specified.  This method is untested and barely works.  Just left it in for anyone else who might want to play with it.

## Setup

To compile the service and build the docker image, enter the following in your terminal starting in the root of this project (assuming you have docker set up and working):
``` bash 

# build the application
cd src/skid-pdf; go build;

# go back to root of the project
cd ../../

# create docker image with all required elements in place
docker build -t skidpdf -f skid-pdf.dockerfile .

# turn the docker image on
docker run -p 8080:8080 -it skidpdf
```

## TODO:

- set up queue system; going to use AMQP
- include configuration file for port selection and other things.