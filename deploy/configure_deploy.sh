#!/bin/bash
# for any error, stop further execution
set -e
# disable the host key checking
chmod +x ./deploy/disable_host_key_checking.sh
./deploy/disable_host_key_checking.sh

DEPLOY_SERVERS=$DEPLOY_SERVERS
# lets split this string and convert this into array
# In UNIX, we can use this commond to do this
# ${string//substring/replacement}
# our substring is "," and we replace it with nothing.
ALL_SERVERS=(${DEPLOY_SERVERS//,/ })

# Make pem file read only, otherwise bad permission error occurs
PEM_FILE_PATH=$HOME/secrets/ec2-ssh-access.pem
chmod 400 $PEM_FILE_PATH

# Iterate over multiple servers(if needed) and run the shell script there
for server in "${ALL_SERVERS[@]}"
do
  echo "Deploying to ${server}"
  ssh -i $PEM_FILE_PATH ubuntu@${server} 'bash -s' < ./deploy/update_and_restart.sh
done