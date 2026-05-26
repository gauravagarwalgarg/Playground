#include <iostream>

using namespace std;

class Amplifier {
  public:
    void on() { cout << "Amplifier is ON" << endl; }
    void off() { cout << "Amplifier is OFF" << endl; }
    void setVolume(int level) {
      cout << "Amplifier volume set to " << level << endl;
    }
};

class DVDPlayer {
  public:
    void on() { cout << "DVD Player is ON" << endl; }
    void off() { cout << "DVD Player is OFF" << endl; }
    void play(string movie) {
      cout << "Playing movie... " << movie << endl;
    }
};

class Projector {
  public:
    void on() { cout << "Projector is ON" << endl; }
    void off() { cout << "Projector is OFF" << endl; }
    void setInputSource(string source) {
      cout << "Projector input source set to" << source << endl;
    }
};

class HomeTheatreFacade {
  private:
    Amplifier amp;
    DVDPlayer dvd;
    Projector projector;

  public:
    void watchMovie(string movie) {
      cout << "Setting up home theater for movie night..." << endl;
      amp.on();
      amp.setVolume(10);
      projector.on();
      projector.setInputSource("DVD");
      dvd.on();
      dvd.play(movie);
      cout << "Enjoy the movie" << endl;
    }

    void endMovie() {
      cout << "Shutting down home theator..." << endl;
      dvd.off();
      projector.off();
      amp.off();
    }
};

int main() {
  HomeTheatreFacade homeTheater;
  homeTheater.watchMovie("inception");
  homeTheater.endMovie();

  return 0;
}

