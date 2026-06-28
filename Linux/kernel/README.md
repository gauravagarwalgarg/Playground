# Linux Kernel

## Subsystems Overview

### Process Scheduler
- **CFS (Completely Fair Scheduler)**: Default scheduler, uses red-black tree
- **Real-time schedulers**: SCHED_FIFO, SCHED_RR
- **Concepts**: Nice values, priority, time slices, runqueues
- **Key files**: `kernel/sched/`

### Memory Management
- **Virtual memory**: Page tables, TLB, address spaces
- **Page allocator**: Buddy system for physical pages
- **Slab allocator**: kmalloc, kmem_cache for kernel objects
- **Swapping**: Page reclaim, LRU lists, OOM killer
- **Key files**: `mm/`

### VFS (Virtual File System)
- **Abstraction layer**: Uniform interface for all filesystems
- **Key structures**: superblock, inode, dentry, file
- **Operations**: open, read, write, close → dispatched to specific FS
- **Filesystems**: ext4, XFS, Btrfs, tmpfs, procfs, sysfs

### Networking Stack
- **Layers**: Socket → Transport (TCP/UDP) → Network (IP) → Link (Ethernet)
- **Key structures**: sk_buff, sock, net_device
- **Netfilter**: iptables/nftables hook points
- **Key files**: `net/`

### Device Drivers
- **Types**: Character, Block, Network
- **Model**: Bus → Driver → Device
- **Device Tree**: Hardware description for ARM/embedded
- **Key APIs**: platform_driver, probe/remove, file_operations
- **Key files**: `drivers/`

---

## Building the Kernel

```bash
make menuconfig    # Configure
make -j$(nproc)    # Build
make modules_install
make install
```

## Key References

- [kernel.org documentation](https://www.kernel.org/doc/html/latest/)
- Linux Device Drivers (LDD3)
- Understanding the Linux Kernel (Bovet & Cesati)
