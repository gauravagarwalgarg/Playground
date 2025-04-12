CC = gcc
CXX = g++
CFLAGS = -Wall -Wextra -O2
CXXFLAGS = -Wall -Wextra -O2
LDFLAGS = 
SRC_DIR = .
BUILD_DIR = build

C_SOURCES := $(wildcard $(SRC_DIR)/*.c)
CPP_SOURCES := $(wildcard $(SRC_DIR)/*.cpp)
C_OBJS := $(patsubst $(SRC_DIR)/%.c, $(BUILD_DIR)/%.o, $(C_SOURCES))
CPP_OBJS := $(patsubst $(SRC_DIR)/%.cpp, $(BUILD_DIR)/%.o, $(CPP_SOURCES))
C_EXES := $(patsubst $(SRC_DIR)/%.c, $(SRC_DIR)/%, $(C_SOURCES))
CPP_EXES := $(patsubst $(SRC_DIR)/%.cpp, $(SRC_DIR)/%, $(CPP_SOURCES))

all: $(C_EXES) $(CPP_EXES)

$(SRC_DIR)/%: $(BUILD_DIR)/%.o
	$(CC) $(LDFLAGS) $< -o $@

$(SRC_DIR)/%: $(BUILD_DIR)/%.o
	$(CXX) $(LDFLAGS) $< -o $@

$(BUILD_DIR)/%.o: $(SRC_DIR)/%.c | $(BUILD_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp | $(BUILD_DIR)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR) $(C_EXES) $(CPP_EXES)

.PHONY: all clean

