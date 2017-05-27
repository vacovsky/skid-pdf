FROM centos
RUN yum install -y epel-release
RUN yum install -y xorg-x11-fonts-75dpi
RUN yum install -y xorg-x11-fonts-Type1
RUN yum install -y wget
RUN yum install -y wkhtmltopdf
# RUN yum install -y bash
ADD src/skid-pdf/skid-pdf /
CMD ["/skid-pdf"]

# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# docker run -p 8080:8080 -it skidpdf
# docker build -t skidpdf -f skid-pdf.dockerfile .