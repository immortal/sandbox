https://stackoverflow.com/a/17955149/1135424

Compile the code: `gcc -o firstdaemon daemonize.c`
Start the daemon: `./firstdaemon`
Check if everything is working properly: `ps -xj | grep firstdaemon`

The output should be similar to this one:

    +------+------+------+------+-----+-------+------+------+------+-----+
    | PPID | PID  | PGID | SID  | TTY | TPGID | STAT | UID  | TIME | CMD |
    +------+------+------+------+-----+-------+------+------+------+-----+
    |    1 | 3387 | 3386 | 3386 | ?   |    -1 | S    | 1000 | 0:00 | ./  |
    +------+------+------+------+-----+-------+------+------+------+-----+

What you should see here is:

* The daemon has no controlling terminal (__TTY = ?__)
* The parent process ID (__PPID__) is __1__ (The init process)
* The __PID != SID__ which means that our process is NOT the session leader (because of the second fork())
* Because PID != SID our process __can't take control of a TTY again__

Reading the syslog:

Locate your syslog file. Mine is here: `/var/log/syslog`
Do a: `grep firstdaemon /var/log/syslog`

The output should be similar to this one:

    firstdaemon[3387]: First daemon started.
    firstdaemon[3387]: First daemon terminated.
