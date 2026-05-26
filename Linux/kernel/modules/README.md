# Linux Kernel Modules

## What is a Kernel Module?
- Loadable code that extends kernel functionality at runtime
- No need to recompile or reboot the kernel
- Examples: device drivers, filesystems, network protocols

## Minimal Module Template

```c
#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Your Name");
MODULE_DESCRIPTION("A simple kernel module");

static int __init hello_init(void) {
    pr_info("Hello, kernel!\n");
    return 0;
}

static void __exit hello_exit(void) {
    pr_info("Goodbye, kernel!\n");
}

module_init(hello_init);
module_exit(hello_exit);
```

## Makefile for Out-of-Tree Module

```makefile
obj-m += hello.o

KDIR := /lib/modules/$(shell uname -r)/build

all:
	make -C $(KDIR) M=$(PWD) modules

clean:
	make -C $(KDIR) M=$(PWD) clean
```

## Commands

```bash
# Build
make

# Load module
sudo insmod hello.ko

# Check messages
dmesg | tail

# List loaded modules
lsmod | grep hello

# Module info
modinfo hello.ko

# Unload
sudo rmmod hello

# Auto-load with dependencies
sudo modprobe hello
```

## Module Parameters

```c
static int count = 1;
module_param(count, int, 0644);
MODULE_PARM_DESC(count, "Number of iterations");
```

```bash
sudo insmod hello.ko count=5
```

## Key Concepts

- `__init` / `__exit`: Memory optimization markers
- `MODULE_LICENSE("GPL")`: Required for accessing GPL-only symbols
- `pr_info`, `pr_err`, `pr_debug`: Kernel logging macros
- `/proc` and `/sys` interfaces for userspace communication
