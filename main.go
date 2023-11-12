package main

import (
	"flag"
	"fmt"
	"github.com/hsmade/esphome-go/pkg/server"
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/hsmade/growatt-esphome/sensors"
	"github.com/hsmade/growatt-sniffer/pkg/decode"
	"github.com/hsmade/growatt-sniffer/pkg/decrypt"
	"github.com/hsmade/growatt-sniffer/pkg/tzsp"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"log"
	"log/slog"
	"net/http"
	"os"
)

var (
	verbose     = flag.Bool("verbose", false, "enable verbose logging")
	port        = flag.Int("port", 5279, "the port to listen on for the incoming tzsp stream")
	metricsPort = flag.Int("metrics-port", 9001, "the port to export prometheus metrics on")
)

func main() {
	flag.Parse()

	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	if *verbose {
		programLevel.Set(slog.LevelDebug)
	}

	// Set up metrics
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%d", *metricsPort), nil)

	sensorDefinitions := sensors.GetSensorDefinitions()
	SensorUpdates := make(chan conf.SensorUpdate, 1)

	S := server.Server{
		Port: 6053,
		Config: conf.Config{
			DeviceInfo: conf.DeviceInfo{
				Name:         "Growatt Inverter",
				UsesPassword: true,
			},
			Updates: SensorUpdates,
			VerifyPasswordCallback: func(s string) bool {
				// password checking logic goes here
				return true
			},
			Sensors: sensorDefinitions,
		},
	}

	// start the ESP home server
	go func() {
		err := S.Listen()
		if err != nil {
			slog.Error("main: server exited: %w", err)
		}
	}()

	// listen on tzsp stream
	packets := make(chan []byte, 1)
	go func() {
		slog.Info("starting tzsp listener")
		err := tzsp.Listen(*port, packets)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// handle the received packets
	for {
		select {
		case packet := <-packets:
			slog.Debug("received packet", "size", len(packet))
			if len(packet) == 0 {
				continue
			}

			if len(packet) < 100 {
				slog.Debug("received packet too small", "size", len(packet))
				continue
			}

			// parse packet data into struct
			data := decode.Data{}
			err := decode.UnmarshalBinary(decrypt.Decrypt(packet), &data)
			if err != nil {
				slog.Warn("failed to handle packet", "error", err.Error())
				continue
			}

			// do something with the data
			fmt.Printf("%+v\n", data)
			SensorUpdates <- conf.TextSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Status"),
				},
				State: data.Status.String(),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Power In"),
				},
				State: float32(data.PowerIn),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV1 Voltage"),
				},
				State: float32(data.PV1Voltage),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV1 Current"),
				},
				State: float32(data.PV1Current),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV1 Power"),
				},
				State: float32(data.PV1Power),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV2 Voltage"),
				},
				State: float32(data.PV2Voltage),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV2 Current"),
				},
				State: float32(data.PV2Current),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV2 Power"),
				},
				State: float32(data.PV2Power),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Total Power Out"),
				},
				State: float32(data.TotalPowerOut),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Frequency"),
				},
				State: float32(data.Frequency),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Grid Fase1 Voltage"),
				},
				State: float32(data.GridFase1Voltage),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Grid Fase1 Current"),
				},
				State: float32(data.GridFase1Current),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Grid Fase1 Power"),
				},
				State: float32(data.GridFase1Power),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Total Energy Today"),
				},
				State: float32(data.EnergyToday),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Total Energy"),
				},
				State: float32(data.EnergyTotal),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV Total Energy"),
				},
				State: float32(data.PVEnergyTotal),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV1 Total Energy Today"),
				},
				State: float32(data.PV1EnergyToday),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV1 Total Energy"),
				},
				State: float32(data.PV1EnergyTotal),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV2 Total Energy Today"),
				},
				State: float32(data.PV2EnergyToday),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt PV2 Total Energy"),
				},
				State: float32(data.PV2EnergyTotal),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Inverter temperature"),
				},
				State: float32(data.InverterTemperature),
			}
			SensorUpdates <- conf.GenericSensorState{
				BaseSensorState: conf.BaseSensorState{
					Key: sensors.GetKey("Growatt Intelligent Power Management Temperature"),
				},
				State: float32(data.IntelligentPowerManagementTemperature),
			}
		}
	}
}
