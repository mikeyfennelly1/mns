BUILD := ./build
SRC := $(wildcard ./src/*.c)
OBJ := $(patsubst %.c,%.o,$(SRC))

all: $(OBJ)

%.o: $(BUILD) %.c
	@gcc -c $< -o $(BUILD)/$@
$(BUILD):
	@mkdir -p ./build
clean:
	@rm -rf $(BUILD)
=
