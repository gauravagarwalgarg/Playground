#include <iostream>
#include <string>

using namespace std;

class ModernAudioPlayer {
  public:
    virtual void playAudio(string filename) = 0;
    virtual ~ModernAudioPlayer() {}
};

// Adaptee
class OldMediaPlayer {
  public:
    void playfile(string filename) {
      cout << "Playing audio file (oldMediaPlayer): " << filename << endl;
    }
};


// Adaptor class
class MediaAdapter : public ModernAudioPlayer {
  private:
    OldMediaPlayer* oldPlayer;

  public:
    MediaAdapter(OldMediaPlayer* oldPlayer) : oldPlayer(oldPlayer) {}

    void playAudio(string filename) override {
      cout << "Adapter converting request to oldMediaPlayer..." << endl;
      oldPlayer->playfile(filename);
    }
};


int main() {
  OldMediaPlayer oldPlayer;

  ModernAudioPlayer* adaptedPlayer = new MediaAdapter(&oldPlayer);

  adaptedPlayer->playAudio("song.mp3");

  delete adaptedPlayer;

  return 0;
}


