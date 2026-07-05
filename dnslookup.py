import sys
# pyrefly: ignore [missing-import]
import dns.resolver

def lookup(domain):
    record_types = ["A", "AAAA", "MX", "NS", "TXT", "CNAME"]
    for rtype in record_types:
        try:
            answers = dns.resolver.resolve(domain, rtype)
            print(f"\n{rtype} Records:")
            for rdata in answers:
                print(f"  {rdata.to_text()}")
        except dns.resolver.NoAnswer:
            continue
        except dns.resolver.NXDOMAIN:
            print(f"Domain '{domain}' does not exist.")
            return
        except Exception as e:
            print(f"  Error fetching {rtype}: {e}")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python dns_lookup.py <domain>")
        sys.exit(1)

    domain = sys.argv[1]
    print(f"DNS Lookup for: {domain}")
    lookup(domain)