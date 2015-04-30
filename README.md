# loopback

Just in case you don't have a spare disk or partition but would like to play around with btrfs this tool creates a loopback device from a file which you can mount whereever you want.

Usecases:
 * systemd machinectl
 * conair (https://github.com/teemow/conair)

# usage

Create a btrfs fs for conair and machinectl

```
sudo loopback create --name=conair --size=10 /var/lib/conair
sudo loopback create --name=machinectl --size=10 /var/lib/machinectl
```

Remove the loopbacks

```
sudo loopback destroy --name=conair
sudo loopback destroy --name=machinectl
```
