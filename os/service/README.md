<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse: `service` Library
**URL** [multiverse-os.org](https://multiverse-os.org)

**Introduction**
A general daemon service toolkit for handling basic Linux services, enabling
daemonizaton of a process, creation of a pid file, interacting with syslog, and
signal handling. 

Future development plans include using the Multiverse log library to fallback if
syslog is not available, and creation of init scripts for both systemd and sysV.

This will most likely become a subcomponent of the Multiverse CLI Go framework.

**Functional Requrements**

  * Create child processes (and general process and child process management) to
    enable rolling restart, no-downtime updates, restart on file change in
    development mode and so on
  * Daemonize 
  * Handle pids
  * Hand syslog, send fmt to syslog instead of screen 
