#!/bin/sh

validate(){

if [ -z "$1" ]; then
    echo "username param required"
    return 1
fi

if [ -z "$2" ]; then
    echo "user email param required"
    return 1
fi

return 0
}

validate "$@" 

if [ $? -eq 1  ]
then 
    exit 1
fi

random=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 16 | head -n 1)

echo "BOOT_ENV=local-$1" > .env
echo "BOOT_DEFAULT_DESTINATION_EMAIL_ADDRESS=\"$2\"" >> .env
echo "BOOT_AUTH_JWT_PRIVATE_KEY=$random" >> .env


