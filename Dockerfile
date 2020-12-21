FROM alpine
# Add Maintainer Info
LABEL maintainer="Duy Ha <duyhph@gmail.com>"
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy exec file and config
COPY main ./
RUN echo "nameserver 8.8.8.8" > /etc/resolv.conf
RUN echo "nameserver 8.8.4.4" > /etc/resolv.conf
# Run the executable
CMD ["./main"]