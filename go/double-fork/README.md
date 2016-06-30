http://stackoverflow.com/a/881434/1135424

# Fork a second child and exit immediately to prevent zombies.  This
# causes the second child process to be orphaned, making the init
# process responsible for its cleanup.  And, since the first child is
# a session leader without a controlling terminal, it's possible for
# it to acquire one by opening a terminal in the future (System V-
# based systems).  This second fork guarantees that the child is no
# longer a session leader, preventing the daemon from ever acquiring
# a controlling terminal.
