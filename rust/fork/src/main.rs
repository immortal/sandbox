use fork::{daemon, Fork};
use std::process::{id, Command};

// ps -axo ppid,pid,pgid,sess,tty,tpgid,stat,uid,user,%mem,%cpu,command, | egrep "fork|sleep|PID"
fn main() {
    if let Ok(Fork::Child) = daemon(false) {
        println!("my pid {}", id());
        Command::new("sleep")
            .arg("300")
            .output()
            .expect("failed to execute process");
    }
}
