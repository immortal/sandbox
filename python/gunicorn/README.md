# install gunicorn

    $ pip install gunicorn

Run:

    $ gunicorn -w 2 test:app

Immortalize it:

    $ immortal -p test.pid gunicorn -p test.pid -w 2 test:app

Test with:

    # Reload a new master with new workers
    kill -s USR2 $PID
    # Graceful stop old workers
    kill -s WINCH $OLDPID
    # Graceful stop old master
    kill -s QUIT $OLDPID
