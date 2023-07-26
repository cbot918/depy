docker container inspect ing-control-plane \
  --format '{{ .NetworkSettings.Networks.kind.IPAddress }}'