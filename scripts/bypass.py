
import json
import subprocess

def run_remote_command(server: str, command: str):
    """Run a command on a remote server via SSH"""
    ssh_command = f"ssh {server} '{command}'"
    result = subprocess.run(ssh_command, shell=True, text=True, capture_output=True)
    if result.returncode != 0:
        raise Exception(f"Command failed on {server}: {ssh_command}\n{result.stderr}")
    return result.stdout


def get_container_pids(server: str, service_name: str):
    """Get PIDs of sidecar proxies for a given microservice on a remote server"""
    containers_output = run_remote_command(
        server,
        "sudo crictl --runtime-endpoint unix:///run/containerd/containerd.sock ps",
    )

    # Adjust the following logic based on your container identification method
    sidecar_containers = []
    for line in containers_output.splitlines()[1:]:
        if service_name in line and "istio-proxy" in line:
            container_id = line.split()[0]
            sidecar_containers.append(container_id)

    pids = []
    for container_id in sidecar_containers:
        inspect_output = run_remote_command(
            server,
            f"sudo crictl --runtime-endpoint unix:///run/containerd/containerd.sock inspect {container_id}",
        )
        pid = json.loads(inspect_output)["info"]["pid"]
        pids.append(pid)

    return pids


def bypass_sidecar(hostname: str, service_name: str, port: str, direction: str):
    """Set up iptables rules for each container on a remote server"""
    pids = get_container_pids(hostname, service_name)
    for pid in pids:
        print(f"Setting up {direction} iptables on server {hostname} with PID {pid}")
        if direction == "S":
            # inbound traffic
            run_remote_command(
                hostname,
                f"sudo nsenter -t {pid} -n iptables -t nat -I PREROUTING 1 -p tcp --dport {port} -j ACCEPT -w",
            )
        elif direction == "C":
            # outbound traffic
            run_remote_command(
                hostname,
                f"sudo nsenter -t {pid} -n iptables -t nat -I ISTIO_OUTPUT 1 -p tcp --dport {port} -j RETURN -w",
            )
        else:
            raise ValueError

bypass_sidecar("h2", "frontend", "8082", "C")
bypass_sidecar("h2", "search", "8082", "S")
bypass_sidecar("h2", "search", "8083", "C")
bypass_sidecar("h2", "geo", "8083", "S")

# bypass_sidecar("h2", "frontend", "8081", "C")
# bypass_sidecar("h2", "ping", "8081", "S")

# bypass_sidecar("h2", "productpage", "8080", "S")
# bypass_sidecar("h2", "productpage", "8083", "C")
# # bypass_sidecar("h2", "details", "8081", "S")
# bypass_sidecar("h2", "reviews", "8083", "S")
# bypass_sidecar("h2", "reviews", "8083", "C")
# # bypass_sidecar("h2", "reviews", "8083", "C")
# # bypass_sidecar("h1", "reviews", "27017", "S")
# bypass_sidecar("h2", "mongodb-reviews", "27017", "S")
