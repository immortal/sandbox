# install gunicorn

    $ pip install gunicorn

Run:

    $ gunicorn -w 2 test:app

Immortalize it:

    $ immortal -p test.pid gunicorn -p test.pid -w 2 test:app
