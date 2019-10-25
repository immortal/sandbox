use fork::{daemon, Fork};
use std::process::{id, Command};

// ps -axo ppid,pid,pgid,sess,tty,tpgid,stat,uid,user,%mem,%cpu,command, | egrep "fork|sleep|PID"
fn main() {
    if let Ok(Fork::Child) = daemon(false, false) {
        // println!("my pid {}", id());
        // can't print no stdout/stderr unles daemon(false, true)
        Command::new("sleep")
            .arg("300")
            .output()
            .expect("failed to execute process");
    }
}
