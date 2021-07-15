# a binary runner for plan9

install python, go, qemu, screen, get a 9front.iso in cwd, set `$nineapiauth` to whatever you want, then run `./startvm_run.sh`
after a minute or two, it will exit, you will have a screen session running 9front qemu in text mode, running a 9goapi binary.
This will be exposed on (host) localhost:8080.
Example usage:
```
> curl -X POST -F 'toexec=@/home/jacob/dev/zig/build/test.6' -F 'auth="bruh"' localhost:8080
Hello World!
```
Where the `toexec` form field is the file to execute on plan9, and the auth field must match up with what you set `$nineapiauth` to.
Then the server will return the output of the file, or if there are any errors, return them.
If something goes wrong, just do `screen -r plan9runner` to get into the vm and you can then debug stuff.