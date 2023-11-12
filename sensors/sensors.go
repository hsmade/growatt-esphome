package sensors

import (
	"github.com/hsmade/esphome-go/protobuf/api"
	"hash/fnv"
)

import (
	"github.com/hsmade/esphome-go/pkg/server/conf"
	"github.com/iancoleman/strcase"
)

func GetKey(name string) uint32 {
	key := fnv.New32()
	_, _ = key.Write([]byte(strcase.ToLowerCamel(name))) // name of sensor, should be camelcase
	return key.Sum32()
}

func GetSensorDefinitions() []conf.Sensor {
	return []conf.Sensor{
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_power_in",
					UniqueId: "growatt_power_in",
					Name:     "Growatt Power In",
					Key:      GetKey("Growatt Power In"),
				},
				DeviceClass:       "power",
				UnitOfMeasurement: "W",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.TextSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_status",
					UniqueId: "growatt_status",
					Name:     "Growatt Status",
					Key:      GetKey("Growatt Status"),
				},
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv1_voltage",
					UniqueId: "growatt_pv1_voltage",
					Name:     "Growatt PV1 Voltage",
					Key:      GetKey("Growatt PV1 Voltage"),
				},
				DeviceClass:       "voltage",
				UnitOfMeasurement: "V",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv1_current",
					UniqueId: "growatt_pv1_current",
					Name:     "Growatt PV1 Current",
					Key:      GetKey("Growatt PV1 Current"),
				},
				DeviceClass:       "current",
				UnitOfMeasurement: "A",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv1_power",
					UniqueId: "growatt_pv1_power",
					Name:     "Growatt PV1 Power",
					Key:      GetKey("Growatt PV1 Power"),
				},
				DeviceClass:       "power",
				UnitOfMeasurement: "W",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv2_voltage",
					UniqueId: "growatt_pv2_voltage",
					Name:     "Growatt PV2 Voltage",
					Key:      GetKey("Growatt PV2 Voltage"),
				},
				DeviceClass:       "voltage",
				UnitOfMeasurement: "V",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv2_current",
					UniqueId: "growatt_pv2_current",
					Name:     "Growatt PV2 Current",
					Key:      GetKey("Growatt PV2 Current"),
				},
				DeviceClass:       "current",
				UnitOfMeasurement: "A",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv2_power",
					UniqueId: "growatt_pv2_power",
					Name:     "Growatt PV2 Power",
					Key:      GetKey("Growatt PV2 Power"),
				},
				DeviceClass:       "power",
				UnitOfMeasurement: "W",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_total_power_out",
					UniqueId: "growatt_total_power_out",
					Name:     "Growatt Total Power Out",
					Key:      GetKey("Growatt Total Power Out"),
				},
				DeviceClass:       "power",
				UnitOfMeasurement: "W",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_frequency",
					UniqueId: "growatt_frequency",
					Name:     "Growatt Frequency",
					Key:      GetKey("Growatt Frequency"),
				},
				DeviceClass:       "frequency",
				UnitOfMeasurement: "Hz",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_grid_fase1_voltage",
					UniqueId: "growatt_grid_fase1_voltage",
					Name:     "Growatt Grid Fase1 Voltage",
					Key:      GetKey("Growatt Grid Fase1 Voltage"),
				},
				DeviceClass:       "voltage",
				UnitOfMeasurement: "V",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_grid_fase1_current",
					UniqueId: "growatt_grid_fase1_current",
					Name:     "Growatt Grid Fase1 Current",
					Key:      GetKey("Growatt Grid Fase1 Current"),
				},
				DeviceClass:       "current",
				UnitOfMeasurement: "A",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_grid_fase1_power",
					UniqueId: "growatt_grid_fase1_power",
					Name:     "Growatt Grid Fase1 Power",
					Key:      GetKey("Growatt Grid Fase1 Power"),
				},
				DeviceClass:       "power",
				UnitOfMeasurement: "W",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_total_energy_today",
					UniqueId: "growatt_total_energy_today",
					Name:     "Growatt Total Energy Today",
					Key:      GetKey("Growatt Total Energy Today"),
				},
				DeviceClass:       "energy",
				UnitOfMeasurement: "KWh",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_total_energy",
					UniqueId: "growatt_total_energy",
					Name:     "Growatt Total Energy",
					Key:      GetKey("Growatt Total Energy"),
				},
				DeviceClass:       "energy",
				UnitOfMeasurement: "KWh",
				AccuracyDecimals:  1,
				StateClass:        api.SensorStateClass_STATE_CLASS_TOTAL_INCREASING,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv_total_energy",
					UniqueId: "growatt_pv_total_energy",
					Name:     "Growatt Solar Total Energy",
					Key:      GetKey("Growatt Solar Total Energy"),
				},
				DeviceClass:       "energy",
				UnitOfMeasurement: "KWh",
				AccuracyDecimals:  1,
				StateClass:        api.SensorStateClass_STATE_CLASS_TOTAL_INCREASING,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv1_total_energy_today",
					UniqueId: "growatt_pv1_total_energy_today",
					Name:     "Growatt PV1 Total Energy Today",
					Key:      GetKey("Growatt PV1 Total Energy Today"),
				},
				DeviceClass:       "energy",
				UnitOfMeasurement: "KWh",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv1_total_energy",
					UniqueId: "growatt_pv1_total_energy",
					Name:     "Growatt PV1 Total Energy",
					Key:      GetKey("Growatt PV1 Total Energy"),
				},
				DeviceClass:       "energy",
				UnitOfMeasurement: "KWh",
				AccuracyDecimals:  1,
				StateClass:        api.SensorStateClass_STATE_CLASS_TOTAL_INCREASING,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv2_total_energy_today",
					UniqueId: "growatt_pv2_total_energy_today",
					Name:     "Growatt PV2 Total Energy Today",
					Key:      GetKey("Growatt PV2 Total Energy Today"),
				},
				DeviceClass:       "energy",
				UnitOfMeasurement: "KWh",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_pv2_total_energy",
					UniqueId: "growatt_pv2_total_energy",
					Name:     "Growatt PV2 Total Energy",
					Key:      GetKey("Growatt PV2 Total Energy"),
				},
				DeviceClass:       "energy",
				UnitOfMeasurement: "KWh",
				AccuracyDecimals:  1,
				StateClass:        api.SensorStateClass_STATE_CLASS_TOTAL_INCREASING,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_inverter_temperature",
					UniqueId: "growatt_inverter_temperature",
					Name:     "Growatt Inverter temperature",
					Key:      GetKey("Growatt Inverter temperature"),
				},
				DeviceClass:       "temperature",
				UnitOfMeasurement: "°C",
				AccuracyDecimals:  1,
			},
		},
		{
			Definition: conf.GenericSensorDefinition{
				BaseSensorDefinition: conf.BaseSensorDefinition{
					ObjectId: "growatt_ipm_temperature",
					UniqueId: "growatt_ipm_temperature",
					Name:     "Growatt Intelligent Power Management Temperature",
					Key:      GetKey("Growatt Intelligent Power Management Temperature"),
				},
				DeviceClass:       "temperature",
				UnitOfMeasurement: "°C",
				AccuracyDecimals:  1,
			},
		},
	}
}
