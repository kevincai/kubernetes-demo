FROM alpine:3.8
COPY bin/gogreeting /bin/gogreeting
EXPOSE 8080/TCP
CMD ["/bin/gogreeting"]
