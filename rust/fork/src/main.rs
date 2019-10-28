use fork::{daemon, fork, Fork};
use std::io;
use std::process::Command;

// ps -axo ppid,pid,pgid,sess,tty,tpgid,stat,uid,user,%mem,%cpu,command, | egrep "fork|sleep|PID"
fn main() {
    if let Ok(Fork::Child) = fork() {
        println!("normal fork");
        Command::new("sleep")
            .arg("500")
            .output()
            .expect("failed to execute process");
    }

    // read from stdin daemon(false, true)
    // if false will not fork
    if let Ok(Fork::Child) = daemon(false, true) {
        let mut input = String::new();
        io::stdin()
            .read_line(&mut input)
            .expect("failed to read line");

        if input.trim() == "fork" {
            println!("wow");

            //if let Ok(Fork::Child) = daemon(false, false) {
            // println!("my pid {}", id());
            // can't print no stdout/stderr unles daemon(false, true)
            Command::new("sleep")
                .arg("300")
                .output()
                .expect("failed to execute process");
        }
    };
}
