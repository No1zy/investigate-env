FROM openjdk:12
COPY dist/java/src/main/ /src/main
WORKDIR ./src/main
RUN javac Main.java
CMD ["java", "Main"]
