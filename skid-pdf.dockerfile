FROM scratch
ADD src/skid-pdf/skid-pdf /
CMD ["/skid-pdf"]
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# docker run -p 8080:8080 -it skidpdf
# docker build -t skidpdf -f skid-pdf.dockerfile .