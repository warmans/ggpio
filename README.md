# GGPIO

"Generic GPIO" interface with multiple adapters.

Writing code for a Pi is annoying because you 
need to basically develop the code on the device itself 
or at least have a process of syncing a binary every
time you want to run the code. The RTk device solves 
this problem to some extent but all the software is 
written in python, and python is stupid.

So this library exposes a generic interface that is implemented 
by adapters for both raspberry pi (via the "rpio" module by stianeikeland) and 
rtk.

Currently only reading and writing from pins is supported, so no
PWM or edge detection or whatever.