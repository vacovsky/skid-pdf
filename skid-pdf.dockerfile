FROM centos
RUN yum install -y epel-release
RUN yum install -y wkhtmltopdf openssl xorg-x11-server-Xvfb urw-fonts
RUN echo -e '#!/bin/bash\nxvfb-run -a --server-args="-screen 0, 1024x768x24" /usr/bin/wkhtmltopdf -q $*' > /usr/bin/wkhtmltopdf.sh; chmod a+x /usr/bin/wkhtmltopdf.sh; ln -s /usr/bin/wkhtmltopdf.sh /usr/local/bin/wkhtmltopdf
ADD src/skid-pdf/skid-pdf /
ADD src/skid-pdf/skidpdf_settings.json /
CMD ["/skid-pdf"]