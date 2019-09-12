FROM perl:latest

WORKDIR /perl
ADD dist/perl /perl
COPY data /per/data
RUN cpan URI

CMD ["perl", "main.pl"]

