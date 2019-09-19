FROM ruby:latest

WORKDIR /usr/src/app
COPY /dist/rails/Gemfile* ./
COPY ./dist/rails/ .

RUN chmod 755 entrypoint.sh \
    && apt-get update \
    && curl -sL https://deb.nodesource.com/setup_12.x | bash - \
    && apt-get install -y --no-install-recommends \
        mysql-client \
        nodejs \
        yarn \ 
    && curl -o- -L https://yarnpkg.com/install.sh | bash \ 
    && rm -rf /var/lib/apt/lists/* \
    && bundle install \
    && export PATH="$HOME/.yarn/bin:$HOME/.config/yarn/global/node_modules/.bin:$PATH" \
    && yarn install --integrity --check-files 

ENTRYPOINT ["/usr/src/app/entrypoint.sh"]

EXPOSE 3000
CMD ["rails", "server", "-p", "3000", "-b", "0.0.0.0"]
