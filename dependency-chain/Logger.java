import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

// Base Logger class with new abstract method
public abstract class Logger {
    public abstract void log(String message);
    
    protected String formatMessage(String message) {
        String timestamp = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
        return String.format("[%s] %s", timestamp, message);
    }
}

// ConsoleLogger subclass modified
public class ConsoleLogger extends Logger {
    @Override
    public void log(String message) {
        System.out.println("Console: " + formatMessage(message)); // Use new method
    }
}

// FileLogger subclass modified
public class FileLogger extends Logger {
    @Override
    public void log(String message) {
        // Write to a file
        System.out.println("File: " + formatMessage(message)); // Use new method
    }
}
