# Linux Device Drivers

## Character Devices

### Overview
- Accessed as a stream of bytes (like a file)
- Examples: `/dev/tty`, `/dev/random`, serial ports
- Implement `file_operations`: open, read, write, release, ioctl

### Minimal Character Device

```c
#include <linux/fs.h>
#include <linux/cdev.h>
#include <linux/uaccess.h>

static struct cdev my_cdev;
static dev_t dev_num;

static ssize_t my_read(struct file *f, char __user *buf, size_t len, loff_t *off) {
    // copy_to_user(buf, kernel_buf, count);
    return 0;
}

static struct file_operations fops = {
    .owner = THIS_MODULE,
    .read = my_read,
};
```

## Platform Drivers

### Overview
- For devices that don't sit on a discoverable bus (I2C, SPI, memory-mapped)
- Matched to devices via compatible string or device tree

### Template

```c
static int my_probe(struct platform_device *pdev) {
    // Initialize hardware, register with subsystem
    return 0;
}

static int my_remove(struct platform_device *pdev) {
    // Cleanup
    return 0;
}

static const struct of_device_id my_of_match[] = {
    { .compatible = "vendor,my-device" },
    { }
};

static struct platform_driver my_driver = {
    .probe = my_probe,
    .remove = my_remove,
    .driver = {
        .name = "my-driver",
        .of_match_table = my_of_match,
    },
};
module_platform_driver(my_driver);
```

## Device Tree (DT)

### What
- Data structure describing hardware topology
- Compiled from `.dts` (source) to `.dtb` (binary)
- Passed to kernel at boot by bootloader

### Example Node

```dts
my_device@0x40000000 {
    compatible = "vendor,my-device";
    reg = <0x40000000 0x1000>;
    interrupts = <0 42 4>;
    clocks = <&clk_controller 3>;
    status = "okay";
};
```

### Key Properties
- `compatible`: Matches driver's `of_device_id`
- `reg`: Memory-mapped registers (address, size)
- `interrupts`: IRQ specification
- `status`: "okay" or "disabled"

## Driver Development Workflow

1. Write driver with probe/remove
2. Add device tree node or platform_device
3. Build as module: `make modules`
4. Load: `insmod` or `modprobe`
5. Verify: `dmesg`, `/sys/class/`, `/dev/`
