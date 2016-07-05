require "rack"

counter = File.read("counter.txt").to_i + 1
File.open("counter.txt", "w") { |f| f.write(counter) }

sleep 2

run(proc do |env|
  [ "200", { "Content-Type" => "text/plain" }, [ counter.to_s ] ]
end)
