This creates a Session ID(SID):

    $ ps  48952
      PID   TT  STAT      TIME COMMAND
      48952   ??  Ss     0:00.00 ./main -daemonize=false

> Notice the Ss

``ps -d 48952`` will not show the procces
