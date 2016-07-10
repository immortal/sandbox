Create counter.txt:

    $ echo 0 > counter.txt

Run:

    bundle exec unicorn -c unicorn.rb

Immortalize it:

    $ immortal -p ./unicorn.pid bundle exec unicorn -c unicorn.rb
