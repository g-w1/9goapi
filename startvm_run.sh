#!/bin/sh

GOOS=plan9 go build 9goapi.go

# curl http://fulton.software/9front/9front-8392.16c5ead832f2.amd64.iso > 9front.iso

screen -wipe

screen -d -m -S plan9runner qemu-system-x86_64 -cpu host -enable-kvm -m 1024 -net nic,model=virtio,macaddr=52:54:00:00:EE:03 -net user -device virtio-scsi-pci,id=scsi -drive if=none,id=vd1,file=9front.iso -device scsi-cd,drive=vd1,bootindex=0 -device e1000,netdev=net0 -netdev user,id=net0,hostfwd=tcp::8080-:8080 -curses
sleep 6
function send() {
	screen -S plan9runner -p 0 -X stuff "$1"
	sleep .5
}
send "^M"
send "^M"
sleep 1
send "text^M"
sleep 13
send "ip/ipconfig^M"
timeout 20 python3 -m http.server 6969 &
send "hget http://10.0.2.2:6969/9goapi > /tmp/9goapi^M"
sleep 1
send "chmod +x /tmp/9goapi^M"
# authenticate
send "9apiauth=$nineapiauth^M"
# lets go!
send "/tmp/9goapi^M"
