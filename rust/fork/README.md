check pid:

    ps -axo ppid,pid,pgid,sess,tty,tpgid,stat,uid,user,command | egrep "fork|PID"
