#include <iostream>
#include <vector>
#include <unordered_map>
#include <memory>
#include <mutex>
#include <ctime>

// Event Class
class Event {
public:
    int eventId;
    std::string title;
    std::time_t startTime;
    std::time_t endTime;
    std::string description;
    std::vector<std::string> attendees;
    
    Event(int id, std::string t, std::time_t start, std::time_t end, std::string desc)
        : eventId(id), title(std::move(t)), startTime(start), endTime(end), description(std::move(desc)) {}
};

// User Class
class User {
public:
    std::string userId;
    std::vector<int> eventIds;
    
    explicit User(std::string id) : userId(std::move(id)) {}
};

// Calendar Service
class CalendarService {
private:
    std::unordered_map<int, std::shared_ptr<Event>> events;
    std::unordered_map<std::string, std::shared_ptr<User>> users;
    std::mutex mtx;
    int eventCounter = 0;
    
public:
    int createEvent(const std::string& title, std::time_t startTime, std::time_t endTime, 
                    const std::string& description, const std::vector<std::string>& attendees) {
        std::lock_guard<std::mutex> lock(mtx);
        int eventId = ++eventCounter;
        auto event = std::make_shared<Event>(eventId, title, startTime, endTime, description);
        event->attendees = attendees;
        events[eventId] = event;
        
        for (const auto& attendee : attendees) {
            if (users.find(attendee) == users.end()) {
                users[attendee] = std::make_shared<User>(attendee);
            }
            users[attendee]->eventIds.push_back(eventId);
        }
        return eventId;
    }
    
    std::shared_ptr<Event> getEvent(int eventId) {
        std::lock_guard<std::mutex> lock(mtx);
        return events.count(eventId) ? events[eventId] : nullptr;
    }
    
    bool deleteEvent(int eventId) {
        std::lock_guard<std::mutex> lock(mtx);
        if (events.find(eventId) == events.end()) return false;
        events.erase(eventId);
        return true;
    }
};

int main() {
    CalendarService calendar;
    
    std::time_t now = std::time(nullptr);
    int eventId = calendar.createEvent("Meeting", now, now + 3600, "Team Sync-up", {"user1", "user2"});
    
    auto event = calendar.getEvent(eventId);
    if (event) {
        std::cout << "Event: " << event->title << " with attendees: ";
        for (const auto& attendee : event->attendees) {
            std::cout << attendee << " ";
        }
        std::cout << std::endl;
    }
    return 0;
}
