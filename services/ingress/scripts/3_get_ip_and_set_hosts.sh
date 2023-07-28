echo $(docker container inspect ing-control-plane --format '{{ .NetworkSettings.Networks.kind.IPAddress }}') hello.yale.com \
>> /etc/hosts

