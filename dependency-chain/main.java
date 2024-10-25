// Base Logger class
public abstract class Logger {
    public abstract void log(String message);
}

// ConsoleLogger subclass
public class ConsoleLogger extends Logger {
    @Override
    public void log(String message) {
        System.out.println("Console: " + message);
    }
}

// FileLogger subclass
public class FileLogger extends Logger {
    @Override
    public void log(String message) {
        // Write to a file
        System.out.println("File: " + message);
    }
}

// Usage
public class Application {
    private Logger logger;

    public Application(Logger logger) {
        this.logger = logger;
    }

    public void doSomething() {
        logger.log("Doing something!");
    }
}

// Main
Application app = new Application(new ConsoleLogger());
app.doSomething();
