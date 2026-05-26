# Linux Userspace

## systemd

### Overview
- Default init system for most modern Linux distributions
- PID 1: First process, manages all other services
- Unit files: Declarative service configuration

### Key Concepts
- **Units**: service, socket, timer, mount, target
- **Targets**: Groups of units (like runlevels): multi-user.target, graphical.target
- **Dependencies**: Wants, Requires, After, Before

### Commands
```bash
systemctl start/stop/restart <service>
systemctl enable/disable <service>
systemctl status <service>
journalctl -u <service> -f    # Follow logs
```

### Unit File Example
```ini
[Unit]
Description=My Application
After=network.target

[Service]
Type=simple
ExecStart=/usr/bin/myapp
Restart=on-failure
User=myapp

[Install]
WantedBy=multi-user.target
```

---

## Init Process

- **SysV init**: Sequential scripts in `/etc/init.d/`, runlevels 0-6
- **systemd**: Parallel startup, dependency-based, socket activation
- **Boot sequence**: BIOS/UEFI → Bootloader → Kernel → init (PID 1) → Services

---

## cgroups (Control Groups)

### What
- Kernel feature to limit, account, and isolate resource usage
- Foundation of containers (Docker, Kubernetes)

### Resources Controlled
- **cpu**: CPU time allocation
- **memory**: Memory limits, OOM behavior
- **io**: Block I/O bandwidth
- **pids**: Process count limits

### cgroups v2
```bash
# Create a cgroup
mkdir /sys/fs/cgroup/mygroup

# Set memory limit (100MB)
echo 104857600 > /sys/fs/cgroup/mygroup/memory.max

# Add process
echo $PID > /sys/fs/cgroup/mygroup/cgroup.procs
```

---

## Namespaces

### What
- Kernel feature providing process isolation
- Each namespace type isolates a different resource

### Types
| Namespace | Isolates | Flag |
|-----------|----------|------|
| PID | Process IDs | CLONE_NEWPID |
| Network | Network stack | CLONE_NEWNET |
| Mount | Filesystem mounts | CLONE_NEWNS |
| UTS | Hostname | CLONE_NEWUTS |
| IPC | IPC resources | CLONE_NEWIPC |
| User | UID/GID mappings | CLONE_NEWUSER |
| Cgroup | Cgroup root | CLONE_NEWCGROUP |

### Commands
```bash
# Create new namespace
unshare --pid --fork --mount-proc bash

# Enter existing namespace
nsenter -t $PID -n    # Enter network namespace of PID

# List namespaces
lsns
```

---

## Containers = cgroups + namespaces + filesystem isolation

Docker/Podman use these primitives to create isolated environments without full VM overhead.
