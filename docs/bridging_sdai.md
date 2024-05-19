# Bridging the gap between Mainnet and Foreign Domains

Nova bridges sDAI (or uses already bridged sDAI) by harnessing the canonical domain bridges, for example:

- The [Optimism Superchain Bridge](https://app.optimism.io/bridge/deposit) for Optimism and BASE
- The [Arbitrum Bridge](https://bridge.arbitrum.io/) for Arbitrum
- The [Polygon PoS bridge](https://portal.polygon.technology/bridge) for Polygon Mainnet

This ensures native compatibility with target domains and reduces trust between the end user and Nova.
The bridged sDAI token is entrusted to the chain's own bridging mechanism.

!!! warning
Integrators and users should keep in mind that funds on a domain other than Ethereum Mainnet are at the mercy of its bridge, no matter which asset they hold.

### Caveats

Bridging in this fashion removes the direct connection between sDAI on foreign domains, and the canonical sDAI contract on Mainnet.
This is due to the fact that redemptions for sDAI on the foreign side cannot communicate through the bridge at this point in time. We solve this by introducing market mechanisms that help keep foreign sDAI at a fair price compared to the canonical price provided by the Mainnet sDAI contract.

## Getting sDAI's price right

> Section about how we ensure deep liquidity for different sDAI-stable pairs on foreign domains

## Supported Venues

> Section about the venues supported for exchange on different foreign domains
