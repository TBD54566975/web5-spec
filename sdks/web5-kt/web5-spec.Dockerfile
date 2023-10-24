# Use an ARM64 compatible OpenJDK image
FROM openjdk:17-jdk

# Set the working directory to /app
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . /app/

# Build the project
RUN ./gradlew build

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./gradlew", "run"]