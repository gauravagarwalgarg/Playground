"""
Facade Pattern

Provides a simplified interface to a complex subsystem. HomeTheaterFacade
wraps TV, Speakers, Lights, and Player into simple watch/stop operations.
"""


class TV:
    def __init__(self):
        self.on = False
        self.input_source = ""

    def turn_on(self) -> str:
        self.on = True
        return "TV on"

    def set_input(self, source: str) -> str:
        self.input_source = source
        return f"TV input: {source}"

    def turn_off(self) -> str:
        self.on = False
        return "TV off"


class Speakers:
    def __init__(self):
        self.volume = 0

    def turn_on(self) -> str:
        return "Speakers on"

    def set_volume(self, level: int) -> str:
        self.volume = level
        return f"Volume: {level}"

    def turn_off(self) -> str:
        self.volume = 0
        return "Speakers off"


class Lights:
    def __init__(self):
        self.brightness = 100

    def dim(self, level: int) -> str:
        self.brightness = level
        return f"Lights dimmed to {level}%"

    def on_full(self) -> str:
        self.brightness = 100
        return "Lights full"


class Player:
    def __init__(self):
        self.playing = ""

    def play(self, movie: str) -> str:
        self.playing = movie
        return f"Playing: {movie}"

    def stop(self) -> str:
        self.playing = ""
        return "Player stopped"


class HomeTheaterFacade:
    def __init__(self):
        self.tv = TV()
        self.speakers = Speakers()
        self.lights = Lights()
        self.player = Player()

    def watch_movie(self, movie: str) -> list[str]:
        return [
            self.lights.dim(20),
            self.tv.turn_on(),
            self.tv.set_input("HDMI1"),
            self.speakers.turn_on(),
            self.speakers.set_volume(8),
            self.player.play(movie),
        ]

    def stop_movie(self) -> list[str]:
        return [
            self.player.stop(),
            self.speakers.turn_off(),
            self.tv.turn_off(),
            self.lights.on_full(),
        ]


if __name__ == "__main__":
    theater = HomeTheaterFacade()

    # One call sets up everything
    steps = theater.watch_movie("Inception")
    assert "Lights dimmed to 20%" in steps
    assert "TV on" in steps
    assert "Playing: Inception" in steps
    assert theater.tv.on is True
    assert theater.speakers.volume == 8
    assert theater.lights.brightness == 20

    # One call tears down everything
    steps = theater.stop_movie()
    assert "Player stopped" in steps
    assert "TV off" in steps
    assert theater.tv.on is False
    assert theater.lights.brightness == 100

    print("All tests passed!")
