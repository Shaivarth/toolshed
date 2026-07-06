import sys
import socket
from concurrent.futures import ThreadPoolExecutor

COMMON_PORTS = {
    21: "FTP", 22: "SSH", 23: "Telnet", 25: "SMTP", 53: "DNS",
    80: "HTTP", 110: "POP3", 143: "IMAP", 443: "HTTPS",
    445: "SMB", 3306: "MySQL", 3389: "RDP", 8080: "HTTP-Alt"
}

def scan_port(target, port, timeout=1.0):
    try:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.settimeout(timeout)
            result = s.connect_ex((target, port))
            if result == 0:
                service = COMMON_PORTS.get(port, "unknown")
                print(f"  Port {port:5} OPEN   ({service})")
    except socket.error:
        pass

def scan(target, start_port, end_port, max_threads=100):
    try:
        ip = socket.gethostbyname(target)
    except socket.gaierror:
        print(f"Could not resolve host: {target}")
        return

    print(f"Scanning {target} ({ip}) ports {start_port}-{end_port}\n")
    with ThreadPoolExecutor(max_workers=max_threads) as executor:
        for port in range(start_port, end_port + 1):
            executor.submit(scan_port, ip, port)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python3 portscanner.py <target> [start_port] [end_port]")
        sys.exit(1)

    target = sys.argv[1]
    start_port = int(sys.argv[2]) if len(sys.argv) > 2 else 1
    end_port = int(sys.argv[3]) if len(sys.argv) > 3 else 1024

    scan(target, start_port, end_port)



# use these commands to run this program
# python3 portscanner.py domain name
# python3 portscanner.py ip address
# python3 portscanner.py domain name start_port end_port
# python3 portscanner.py ip address start_port end_port

    