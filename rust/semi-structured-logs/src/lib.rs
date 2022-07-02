#![allow(unused)]

#[derive(Clone, PartialEq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
}
pub fn log(level: LogLevel, message: &str) -> String {
    let level_str = match level {
        LogLevel::Info => "INFO",
        LogLevel::Warning => "WARNING",
        LogLevel::Error => "ERROR",
    };

    return format!("[{}]: {}", level_str, message);
}
pub fn info(message: &str) -> String {
    log(LogLevel::Info, message)
}
pub fn warn(message: &str) -> String {
     log(LogLevel::Warning, message)
}
pub fn error(message: &str) -> String {
     log(LogLevel::Error, message)
}