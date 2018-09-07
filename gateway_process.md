# Gateway Process

The following is the high level overview of how the packet_forwarder is setup.
From this I can make the LoRa module re-broadcast any received packet, thus turning it into a repeater.

## Control flow

```
        +==============+
        | pktfwd.Run() |
        +==============+
               |
               v
+================================+
| CreateNetworkClient(ttnConfig) |
|  - Creates a client instance   |
+================================+
               |
               v
     +==================+
     | configureBoard() |
     +==================+
               |
               v
       +==============+
       |  Run Manager |
       +==============+
```

## Board Configuration

In `wrapper` `SetBoardConf()` creates a C struct `lgw_conf_board_s` with 2 properties: `clksrc` and `lorawan_public`.
It then calls `lgw_board_setconf()` with the struct and checks for errors through return comparison with `LGW_HAL_SUCCESS`.
For configurations, see `global_conf.json`.

## Manager run

`Manager.run()` calls `wrapper.StartLoRaGateway()` which is a wraps `lgw_start()`.
Control flow then moves to `Manager.handler()` which starts `Manager.startRoutines()` as a go routine.
That starts 3 sub routines, `Manager.uplinkRoutine()`, `Manager.downlinkRoutine()`, and `Manager.startRoutine()`, with an optional gps routine `Manager.gpsRoutine()`.

### uplinkRoutine()
So far this is the only routine of interest in since the repeater is for uplink packets.
It takes a context and an error channel and calls `wrapper.Receive()`.
The packets returns are then validated with `wrapUplinkPayload` and sent back to TTN.


### wrapper.Receive()
Wrapper around `lgw_receive()` which returns arrays of `lgw_pkt_rx_s` structs.

## Implementing

### Receiving LoRa packets
