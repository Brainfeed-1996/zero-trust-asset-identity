# ADR 001: SPIFFE-style Zero-Trust Machine Identity

## Status
Accepted

## Context
Industrial edge assets require a verifiable, cryptographically secure identity to participate in the security mesh. Relying on IP addresses or shared secrets is insufficient for a zero-trust architecture.

## Decision
Implement a microservice that issues SVID-like (SPIFFE Verifiable Identity Document) tokens. 
- Use JWT as the token format for ease of integration with cloud-native tools.
- Identity format: `spiffe://example.org/asset/{asset_id}`.
- Modular storage backend to allow future integration with TPMs or HSMs.

## Consequences
- Every service request must now include an `X-Asset-SVID` header.
- Centralized identity service becomes a critical component (requires high availability).
