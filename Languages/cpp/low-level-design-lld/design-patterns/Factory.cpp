#include <iostream>
using namespace std;

class Logger {
  public:
    virtual void log(const string &message) = 0;
    virtual ~Logger() {}
};

class ConsoleLogger : public Logger {
  public:
    void log(const string &message) override {
      cout << "[Console] " << message << endl;
    }
};

class FileLogger : public Logger {
  public:
    void log(const string& message) override {
      cout << "[File] " << message << endl;
    }
};

class LoggerFactory {
  public:
    virtual Logger* createLogger() = 0;
    virtual ~LoggerFactory() {}
};

class ConsoleLoggerFactory : public LoggerFactory {
  public:
    Logger* createLogger() override {
      return new ConsoleLogger();
    }
};

class FileLoggerFactory : public LoggerFactory {
  public:
    Logger* createLogger() override {
      return new FileLogger();
    }
};

int main() {
  LoggerFactory* factory = nullptr;

  factory = new ConsoleLoggerFactory();
  Logger* logger1 = factory->createLogger();
  logger1->log("Hello! Factory logger console!");
  delete logger1;
  delete factory;

  factory = new FileLoggerFactory();
  Logger* logger2 = factory->createLogger();
  logger2->log("Logging to file...");
  delete logger2;
  delete factory;

  return 0;
}
