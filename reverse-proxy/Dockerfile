FROM nginx

# COPY reverse_proxy_config.txt /etc/nginx/conf.d/reverse-proxy.conf
RUN mv /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.backup
COPY nginx.conf.template /

# RUN /etc/init.d/nginx restart

# CMD ["nginx", "-g", "daemon off;"]

CMD ["/bin/sh" , "-c" , "export GATEWAY_IP=`ip route | awk '/default/ { print $3 }'`  && envsubst '$GATEWAY_IP' < /nginx.conf.template > /etc/nginx/conf.d/reverse-proxy.conf && exec nginx -g 'daemon off;'"]

# mv /etc/nginx/sites-available/reverse-proxy /etc/nginx/conf.d/reverse-proxy.conf
# mv /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.backup