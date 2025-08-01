Objective: Develop a basic logging library that can be used by applications to log messages. The library should handle message logging efficiently and reliably, offering basic configuration options.

Requirements:
    - Multiple Sinks: STDOOUT, FILE, S3
        - sink_type
        - log_level
        - ts_format
    - Multiple Logger: SYNC, ASYNC, CONCURRENT
        - SYNC: Write to all the sinks and then proceed to another message (no concurrency)
        - ASYNC: Populate the channel and procced to next message: Buffer size
            - If buffer overflows, then send suitable error message
        - CONCURRENT: Write using goroutines, mutex locks
    - Ordered
    - No loss
    - Log level order: DEBUG < INFO < WARN < ERROR < FATAL 

SinkType:
    - STDOUT
    - FILE
    - S3
Loglevel:
    - DEBUG
    - INFO
    - WARN
    - ERROR
    - FATAL
TsFormat:
    - DD-MM-YYYY-HH-MM-SS
    - YYYY-DD-MM-HH-MM-SS
    - DD-MM-YYYY
    - HH-MM-SS
Sink:
    - sink_type: SinkType
    - log_level: Loglevel
    - ts_format: TsFormat
    

