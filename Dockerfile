FROM alpine:latest

COPY user-srv /home/

STOPSIGNAL SIGTERM

# CMD ["nginx", "-g", "daemon off;"]
WORKDIR /home/
CMD ./user-srv