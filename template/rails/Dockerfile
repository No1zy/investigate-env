FROM ruby:latest

RUN apt-get update \
    && curl -sL https://deb.nodesource.com/setup_12.x | bash - \
    && apt-get install -y --no-install-recommends \
        mysql-client \
        nodejs \
        yarn \ 
    && curl -o- -L https://yarnpkg.com/install.sh | bash \ 
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app
COPY /dist/rails/Gemfile* ./
COPY ./dist/rails/ .
RUN bundle install

EXPOSE 3000
CMD ["rails", "server", "-p", "3000", "-b", "0.0.0.0"]
