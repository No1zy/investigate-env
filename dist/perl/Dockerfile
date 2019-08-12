FROM perl:latest

WORKDIR /perl
ADD dist/perl /perl
RUN cpan URI

CMD ["perl", "main.pl"]

