# scripts
Useful scripts and bits



### For Galen

Useful things to run

```
cd node-startup

pv-burn # run this script to remove state from any prevous run

docker-compose up  # this will start a functioning node

docker-compose down # this will stop it (will need to run from a new terminal)

# this command will deposit funds for VOTE and tDAI into the account with the pubkey of the wallet
pv-faucet --pubkey e3c154cbf037f6439ebb290e2a45831c5a63d124d0bd24dd3e957504859a8555

# sometimes you have to run this afterward to flush the request through core
pv-timeforward -d 10s

# the above command moves time forward manually. With the nbull blockchain time stays still unless you
# interact by either sending requests to it (like market proposals) or by using the time forward command
```

test

sfdssdfsd