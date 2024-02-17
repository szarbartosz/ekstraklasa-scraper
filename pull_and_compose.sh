#! /bin/bash
echo '
          |          |
       __ |  __   __ | _  __   _
      /  \| /  \ /   |/  / _\ |
      \__/| \__/ \__ |\_ \__  |

                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""\___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
         \    \        __/
          \____\______/

'

echo 'setting up environment variables...'
echo 'PORT=8080' > .env
echo 'GIN_MODE=debug' >> .env
echo 'EUROSPORT_URL=https://eurosport.tvn24.pl/pilka-nozna/pko-bp-ekstraklasa/2023-2024/tabela.shtml' >> .env
echo 'EKSTRAKLASA_URL=https://www.ekstraklasa.org/tabela/sezon' >> .env

echo 'pulling and starting docker container...'
docker compose pull
docker compose up -d