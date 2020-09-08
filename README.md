# Gflog
Print google cloud function structured logs with severity level

## Usage
```
gflog.Print(gflog.Error, "This is message with error severity level error")
gflog.Printf(gflog.Error, "This is message with error severity level %s", gflog.Error)

// Fatal calls Print and os.Exit(1)
gflog.Fatal("This is message with error severity level emergency")
gflog.Fatalf("This is message with error severity level %s", gflog.Emergency)
```

### Available Severity Levels
```
Default
Debug
Info
Notice
Warning
Error
Critical
Alert
Emergency
```
