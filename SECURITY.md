# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.x.x   | :white_check_mark: |

## Reporting a Vulnerability

We take security seriously. If you discover a security vulnerability in sreq, please report it responsibly.

### How to Report

1. **Do NOT** open a public GitHub issue for security vulnerabilities
2. Email the maintainer directly at **mailpriyanshugarg@gmail.com**
3. Include as much detail as possible:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if any)

### What to Expect

- **Acknowledgment**: Within 48 hours of your report
- **Status Update**: Within 7 days with an assessment
- **Resolution**: Timeline depends on severity and complexity

### Severity Levels

| Level    | Description                                    | Response Time |
| -------- | ---------------------------------------------- | ------------- |
| Critical | Remote code execution, credential exposure     | 24-48 hours   |
| High     | Significant data exposure, auth bypass         | 3-5 days      |
| Medium   | Limited data exposure, denial of service       | 1-2 weeks     |
| Low      | Minor issues with limited impact               | Next release  |

## Security Best Practices for Users

When using sreq, follow these security practices:

### Credential Management

- **Never commit** your `.sreq.yaml` config if it contains hardcoded secrets
- Use environment variables or secret managers for sensitive data
- Rotate credentials regularly

### AWS Credentials

- Use IAM roles when possible (especially in CI/CD or EC2)
- Follow the principle of least privilege
- Enable MFA for AWS accounts
- Use separate credentials for development and production

### Consul Tokens

- Use ACL tokens with minimal required permissions
- Rotate Consul tokens periodically
- Never share tokens across environments

### Configuration Files

- Add `*.local.yaml` to your `.gitignore`
- Keep production configs separate from development
- Review configs before sharing or committing

## Security Features in sreq

- **No credential logging**: sreq never logs sensitive values
- **Secure defaults**: Timeouts and connection limits are set conservatively
- **Provider isolation**: Each provider handles its own authentication

## Acknowledgments

We appreciate responsible disclosure and will acknowledge security researchers who help improve sreq's security.
