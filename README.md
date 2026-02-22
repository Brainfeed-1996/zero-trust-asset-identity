# 🛡️ Zero Trust Asset Identity

[![Go](https://img.shields.io/badge/Language-Go-00ADD8.svg)](https://golang.org)
[![Security](https://img.shields.io/badge/Security-Zero--Trust-brightgreen.svg)]()
[![SPIFFE](https://img.shields.io/badge/Standard-SPIFFE-blue.svg)](https://spiffe.io)

Microservice for machine identity issuing SPIFFE-style SVID tokens.

## Architecture
- **API**: Go-based REST API for token issuance and validation.
- **Identity Provider**: Modular backend for identity management.

## SRE/Monitoring
- Metrics endpoint at `/metrics` exporting token lifecycle data.
- Structured logging for security audit trails.

## ADR
- [ADR-001: SPIFFE/SPIRE Alignment for Identity](docs/adr/001-spiffe-alignment.md)
- [ADR-002: JWT for Short-Lived SVIDs](docs/adr/002-jwt-svids.md)
