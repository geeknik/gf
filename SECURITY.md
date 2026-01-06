# Security Policy

## Supported Versions

Only the latest version of `gf` is supported. Users are encouraged to keep their installation up to date.

## Reporting a Vulnerability

If you discover a security vulnerability, please send an email to the maintainer.

The maintainer of this project is **Brian Carpenter** (geeknik).

### What to Include

Please include as much of the following information in your report as possible:

- A description of the vulnerability
- Steps to reproduce the vulnerability
- Potential impact of the vulnerability
- Any suggested mitigation or fix (if known)

### Response Timeline

The maintainer will acknowledge receipt of your report within 7 days and provide regular updates on the status of the issue.

### Disclosure Policy

Once a vulnerability has been fixed, a new release will be pushed as soon as possible. Security advisories will be published along with the fix.

## Best Practices

- `gf` executes external commands (grep, ag, etc.) with user-provided patterns
- Pattern files from untrusted sources should be reviewed before use
- The `-save` flag writes to your home directory - ensure you trust the pattern being saved

## Dependencies

This project uses only the Go standard library and has no external dependencies, minimizing the attack surface.
