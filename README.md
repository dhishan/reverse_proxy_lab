# Reverse Proxy Lab

This project is a quick demo of how reverse proxy works in 2 containers leveraging docker network

How to run
1. Build Image
   ```sh
   $ docker build -t nginx-reverse-proxy:latest . # run this in the reverse proxy folder
   $ docker build -t webapp:latest . # Run this in the webapp folder
   ```
1. Run Container
   ```sh
   $ docker run -p 8080:80 nginx-reverse-proxy:latest # run this in the reverse proxy folder
   $ docker run -p 8000:8000 sample-go-webapp # Run this in the webapp folder
   ```
1. Test & TShoot
   ```sh
   $ docker network inspect bridge # get the docker network info and the gateway ip
   $ docker run --rm -it nicolaka/netshoot # toolkit public container for tshooting
   # Analyse the outputs of the commands
   $ ifconfig
   $ ip route
   $ export GATEWAY_IP=`ip route | awk '/default/ { print $3 }'`
   ```

## References:
1. [Build reverse proxy server in Go](https://dev.to/b0r/implement-reverse-proxy-in-gogolang-2cp4#s1)
2. [How To Set Up a Reverse Proxy Kinsta](https://kinsta.com/blog/reverse-proxy/)