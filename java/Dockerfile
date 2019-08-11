FROM openjdk:12
COPY ./java/src/main/java /src/main
WORKDIR ./src/main
RUN javac Main.java
CMD ["java", "Main"]
