# Changes whalesay to speak a fortune
FROM docker/whalesa:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay
