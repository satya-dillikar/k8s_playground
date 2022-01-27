 #!/usr/bin/env bash

 set -euxo pipefail

 if [[ ! -d $VOLUME ]]; then
   /bin/echo ERR: cannot find the volume $VOLUME, exiting
   exit 1
 fi

 while true; do
   /usr/bin/touch ${VOLUME}/$(hostname)-$(date +%Y%m%d-%H%M%S)
   /bin/sleep 10
 done
