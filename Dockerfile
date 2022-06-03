FROM alpine
ADD helloworld_service /helloworld_service
ENTRYPOINT [ "/helloworld_service" ]
