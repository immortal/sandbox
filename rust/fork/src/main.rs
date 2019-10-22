use nix::unistd::{fork, setsid, ForkResult};
use std::process;

fn main() {
    if let Ok(ForkResult::Child) = fork() {
        let pid = setsid().expect("sesid failed");
        if let Ok(ForkResult::Child) = fork() {
            println!("PGID: {}, my pid {}", pid, process::id());
            process::Command::new("sleep")
                .arg("300")
                .output()
                .expect("failed to execute process");
        }
    }
}
