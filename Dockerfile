FROM centos
RUN yum install -y epel-release
RUN yum install -y http-parser
RUN yum install -y wkhtmltopdf openssl xorg-x11-server-Xvfb urw-fonts npm git
RUN npm install -g bower
RUN cd 
RUN echo -e '#!/bin/bash\nxvfb-run -a --server-args="-screen 0, 1024x768x24" /usr/bin/wkhtmltopdf -q $*' > /usr/bin/wkhtmltopdf.sh; chmod a+x /usr/bin/wkhtmltopdf.sh; ln -s /usr/bin/wkhtmltopdf.sh /usr/local/bin/wkhtmltopdf
RUN mkdir /skidpdf 
ADD bin/skidpdf_x64 /skidpdf/skid-pdf
ADD src/skid-pdf/skidpdf_settings.json /skidpdf
ADD src/skid-pdf/static /skidpdf/static
ADD src/skid-pdf/templates /skidpdf/templates
RUN cd /skidpdf/static; bower install --allow-root --force
RUN groupadd -r skidpdf && useradd -r -g skidpdf skidpdf
CMD cd /skidpdf/; su skidpdf -c "./skid-pdf"
