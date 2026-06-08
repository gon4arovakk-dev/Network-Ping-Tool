import subprocess
import sys

def ping(host):
    try:
        result = subprocess.run(
            ['ping', '-c', '4', host],
            capture_output=True,
            text=True
        )
        print(result.stdout)
    except Exception as e:
        print(f"Error: {e}")

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python ping.py <host>")
        sys.exit(1)
    ping(sys.argv[1])
