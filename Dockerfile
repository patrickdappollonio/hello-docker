FROM scratch
EXPOSE 8000
ADD hello-docker /
ENTRYPOINT ["/hello-docker"]
