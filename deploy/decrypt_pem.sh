#!/bin/bash
# Decrypt the file
mkdir $HOME/secrets

# --batch to prevent interactive command
# --yes to assume "yes" for questions
gpg --quiet --batch --yes --decrypt --passphrase="$SSH_PEM_DECRYPT_KEY" \
--output $HOME/secrets/ec2-ssh-access.pem $PWD/deploy/ec2-ssh-access.pem.gpg