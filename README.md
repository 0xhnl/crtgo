# Certificate Fingerprint Tool (crtgo)

crtgo is a powerful tool designed for passive subdomain enumeration by fetching certificate fingerprints from crt.sh. With crtgo, you can swiftly gather certificate fingerprints for a given domain, aiding in reconnaissance and identifying potential subdomains associated with the target domain.

# Features

Passive Subdomain Enumeration: Fetch certificate fingerprints from crt.sh for a specified domain.
Streamlined Usage: Simple command-line syntax for seamless operation.
Flexible Output: Easily export and analyze the fetched certificate fingerprints.

# Install

You can conveniently install crtgo using the following command:

```bash
go install -v github.com/0xhnl/crtgo@latest
```

# Usage

For single domain:

```bash
$ crtgo -d hackerone.com
design.hackerone.com
docs.hackerone.com
events.hackerone.com
gslink.hackerone.com
hackerone.com
info.hackerone.com
links.hackerone.com
mta-sts.forwarding.hackerone.com
mta-sts.hackerone.com
mta-sts.managed.hackerone.com
support.hackerone.com
www.hackerone.com
```

For a list fo domain:

```bash
$ cat domain.txt
hackerone.com
bugcrowd.com
$ crtgo -i domain.txt
design.hackerone.com
docs.hackerone.com
events.hackerone.com
gslink.hackerone.com
hackerone.com
info.hackerone.com
links.hackerone.com
mta-sts.forwarding.hackerone.com
mta-sts.hackerone.com
mta-sts.managed.hackerone.com
support.hackerone.com
www.hackerone.comapi.bugcrowd.com
blog.bugcrowd.com
bugcrowd.com
collateral.bugcrowd.com
docs.bugcrowd.com
documentation.bugcrowd.com
email.assetinventory.bugcrowd.com
email.bugs.bugcrowd.com
email.crowdcontrol.bugcrowd.com
email.forum.bugcrowd.com
email.submit.bugcrowd.com
events.bugcrowd.com
```
