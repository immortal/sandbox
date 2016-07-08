worker_processes 2
timeout 5
preload_app true
pid "./unicorn.pid"

before_fork do |server, worker|
end


after_fork do |server, worker|
    end
