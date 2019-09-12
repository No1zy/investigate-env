FROM openjdk:12
COPY dist/java/src/main/ /src/main
COPY data/ /src/main/data
WORKDIR ./src/main
RUN javac Main.java
CMD ["java", "Main"]
