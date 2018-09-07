package main

import (
	"context"
	"fmt"
	"github.com/NaNkeen/packet_repeater/wrapper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const initUplinkPollingRate = 100 * time.Microsecond

func main() {
	// System signals
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGABRT)

	// ==================
	// Setup process
	// ==================
	if err := wrapper.SetBoardConf(1, true); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("SX1301 board configured successfully")

	// Configure TX Gain Lut
	if err := wrapper.SetTXGainConf(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("TX Gain Lut configured successfully")

	// Configure RF and SF channels
	if err := wrapper.SetRFChannels(); err != nil {
		fmt.Println(os.Stderr, err)
		return
	}

	if err := wrapper.SetSFChannels(); err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	fmt.Println("RF and SF configured successfully")

	// Configure individual LoRa standard and FSK channels
	if err := configureIndividualChannels(); err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	fmt.Println("LoRa std and FSK channel configured successfully")

	// Start LoRa gateway
	if err := wrapper.StartLoRaGateway(); err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	fmt.Println("LoRa gateway started successfully")

	// TODO Spawn uplink handler Go routines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errc := make(chan error)
	pktc := make(chan wrapper.Packet)

	go uplinkRoutine(ctx, errc, pktc)
	go broadcastRoutine(ctx, errc, pktc)

	select {
	case err := <-errc:
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case sig := <-sigc:
		fmt.Fprintf(os.Stderr, "[!] Signal %s", sig.String())
	}
}

func configureIndividualChannels() error {
	// Configuring LoRa standard channel
	if lora := wrapper.GetLoraSTDChannel(); lora != nil {
		err := wrapper.SetStandardChannel(*lora)
		if err != nil {
			return err
		}
	} else {
		fmt.Fprintln(os.Stderr, "No configuration for LoRa standard channel, ignoring")
	}

	// Configuring FSK channel
	if fsk := wrapper.GetFSKChannel(); fsk != nil {
		err := wrapper.SetFSKChannel(*fsk)
		if err != nil {
			return err
		}
	} else {
		fmt.Fprintln(os.Stderr, "No configuration for FSK standard channel, ignoring")
	}

	return nil
}

func uplinkRoutine(ctx context.Context, errc chan error, pktc chan wrapper.Packet) {
	fmt.Println("Awaiting uplink packets")
	for {
		packets, err := wrapper.Receive()
		if err != nil {
			errc <- err
			return
		}

		if len(packets) == 0 {
			time.Sleep(initUplinkPollingRate)
			continue
		}

		fmt.Printf("Received: %+v\n", packets)

		for _, pkt := range packets {
			pktc <- pkt
		}

		// Check for cancel from context
		select {
		case <-ctx.Done():
			errc <- nil
			return
		default:
			continue
		}
	}
}

func broadcastRoutine(ctx context.Context, errc chan error, pktc chan wrapper.Packet) {
	fmt.Println("Waiting to repeat")
	for {
		select {
		case pkt := <-pktc:
			fmt.Printf("Broadcasting: %+v\n", pkt)
		case <-ctx.Done():
			errc <- nil
			return
		default:
			continue
		}
	}
}
