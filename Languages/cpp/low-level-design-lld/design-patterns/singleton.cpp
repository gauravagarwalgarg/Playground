/**
 * Design Pattern: Singleton (Thread-Safe, Meyer's)
 *
 * Ensures a class has only one instance and provides global access.
 * C++11 guarantees thread-safe initialization of function-local statics.
 */
#include <iostream>
#include <cassert>
using namespace std;

class Logger {
public:
    static Logger& instance() {
        static Logger inst;
        return inst;
    }

    void log(const string& msg) {
        cout << "[LOG] " << msg << endl;
        logCount_++;
    }

    int getLogCount() const { return logCount_; }

    // Delete copy/move
    Logger(const Logger&) = delete;
    Logger& operator=(const Logger&) = delete;

private:
    Logger() : logCount_(0) {}
    int logCount_;
};

int main() {
    Logger& l1 = Logger::instance();
    Logger& l2 = Logger::instance();

    assert(&l1 == &l2); // Same instance

    l1.log("First message");
    l2.log("Second message");
    assert(l1.getLogCount() == 2);

    cout << "All tests passed!" << endl;
    return 0;
}
