Create counter.txt:

    $ echo 0 > counter.txt

Run:

    bundle exec unicorn -c unicorn.rb

Immortalize it:

    $ immortal bundle exec unicorn -c unicorn.rb

Test:

    $ ./immortal -l /tmp/test.log -logger "logger -t unicorn" bundle exec unicorn -c unicorn.rb

Follow pid:

    $ ./immortal -l /tmp/test.log -logger "logger -t unicorn" -f ./unicorn.pid  bundle exec unicorn -c unicorn.rb

Tee:

    $ immortal -l /tmp/test.log -logger "tee /tmp/i.log" -f ./unicorn.pid  bundle exec unicorn -c unicorn.rb

Watch:

    $ watch -n 0.1 "pstree -s unicorn\.rb"
