use libc;
use std::process::{exit, id, Command};

/*
 * The parent forks the child
 * The parent exits
 * The child calls setsid() to start a new session with no controlling terminals
 * The child forks a grandchild
 * The child exits
 * The grandchild is now the daemon
 */

pub enum DaemonError {
    Fork,
}

impl DaemonError {
    fn __description(&self) -> &str {
        match *self {
            DaemonError::Fork => "unable to fork",
        }
    }
}

fn fork() -> Result<(), DaemonError> {
    unsafe {
        let pid = libc::fork();
        if pid < 0 {
            Err(DaemonError::Fork)
        } else if pid == 0 {
            Ok(())
        } else {
            exit(0);
        }
    }
}

fn main() {
    if let Ok(_) = fork() {
        println!("My pid is {}", id());
        unsafe {
            libc::setsid();
        }
        if let Ok(_) = fork() {
            println!("My pid is {}", id());
            Command::new("sleep")
                .arg("300")
                .output()
                .expect("failed to execute process");
        }
    }
}
