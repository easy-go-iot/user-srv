FROM alpine:latest

COPY user-srv /home
COPY grpc_health_probe /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe
STOPSIGNAL SIGTERM

# CMD ["nginx", "-g", "daemon off;"]
WORKDIR /home/
CMD ./user-srv