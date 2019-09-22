FROM openjdk:12

WORKDIR /usr/src/app
COPY dist/spring/ /usr/src/app
COPY data/ /usr/src/app/data
RUN ./gradlew build

CMD ["./gradlew", "bootRun"]
