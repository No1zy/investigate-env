FROM php:7.3-apache

WORKDIR /var/www/html
ADD dist/php-apache /var/www/html/
COPY data /var/www/html/data
COPY dist/php-apache/apache2.conf /etc/apache2/apache2.conf
COPY dist/php-apache/default.conf /etc/apache2/sites-available/000-default.conf
COPY dist/php-apache/ports.conf /etc/apache2/ports.conf

RUN chown www-data -R /var/www/html

EXPOSE 8001

