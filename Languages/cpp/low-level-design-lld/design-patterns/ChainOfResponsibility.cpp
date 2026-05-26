// The Chain of Responsibility (CoR) pattern is a behavioral design pattern that allows multiple objects to handle a request sequentially until one of them processes it. This helps in reducing tight coupling between sender and receiver.
//

#include <iostream>
#include <memory>

using namespace std;

class Logger {
  protected:
    int level;
    shared_ptr<Logger> nextLogger;

  public:
    Logger(int lvl) : level(lvl), nextLogger(nullptr) {}

    void setNextLogger(shared_ptr<Logger> next) {
      nextLogger = next;
    }

    void logMessage(int severity, const string& message) {
      if(severity >= level) {
        write(message);
      }

      if(nextLogger) {
        nextLogger->logMessage(severity, message);
      }
    }

    virtual void write(const string& message) = 0;
};

class InfoLogger : public Logger {
  public:
    InfoLogger(int lvl) : Logger(lvl) {}

    void write(const string& message) override {
      cout << "[INFO]: " << message << endl;
    }
};

class WarningLogger : public Logger {
  public:
    WarningLogger(int lvl) : Logger(lvl) {}

    void write(const string& message) override {
      cout << "[WARNING]: " << message << endl;
    }
};

class ErrorLogger : public Logger {
  public:
    ErrorLogger(int lvl) : Logger(lvl) {}

    void write(const string& message) override {
      cout << "[ERROR]: " << message << endl;
    }
};

shared_ptr<Logger> getLoggerChain() {
  shared_ptr<Logger> errorLogger = make_shared<ErrorLogger>(3);
  shared_ptr<Logger> warningLogger = make_shared<WarningLogger>(2);
  shared_ptr<Logger> infoLogger = make_shared<InfoLogger>(1);

  infoLogger->setNextLogger(warningLogger);
  warningLogger->setNextLogger(errorLogger);

  return infoLogger;
}

int main() {
  auto loggerChain = getLoggerChain();

  loggerChain->logMessage(1, "This is info");
  loggerChain->logMessage(2, "This is warning...");
  loggerChain->logMessage(3, "This is error...");

  return 0;
}

