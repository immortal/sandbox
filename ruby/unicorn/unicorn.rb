worker_processes 2
timeout 5
preload_app true
pid "./unicorn.pid"

before_fork do |server, worker|
    ActiveRecord::Base.connection.disconnect! if defined?(ActiveRecord::Base)
    Sequel::Model.db.disconnect if defined?(Sequel::Model)

    # During a restart, kills old unicorn master after the new master started
    old_pid = "./unicorn.pid.oldbin"
    if File.exists?(old_pid) && old_pid != server.pid
        begin
            # Stop workers of the old master one by one to prevent memory bursts during deployment.
            signal = (worker.nr + 1) >= server.worker_processes ? :QUIT : :TTOU
            Process.kill(signal, File.read(old_pid).to_i)
        rescue Errno::ENOENT, Errno::ESRCH; end
    end

    sleep 3
end


after_fork do |server, worker|
    end
