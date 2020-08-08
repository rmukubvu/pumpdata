package model

const (
	PumpSensorDataByCompany = `
								select
									t.serial_number as SerialNumber,
									p.nick_name as PumpName,
									pt.name as PumpType,
									t.type_text as SensorName,
									t.s_value as SensorValue,
									t.created_date as CreatedDate
								from sensor_data t
								inner join pump p
									on t.serial_number = p.serial_number
								inner join pump_company pc
									on p.id = pc.pump_id
								inner join pump_type pt
									on p.type_id = pt.id
								where pc.company_id = $1
								  and t.created_date >= $2
								  and t.created_date <= $3 `

	AlarmPerPump = `select 	ps.serial_number as SerialNumber,
							st.name as SensorName,
							p.nick_name as PumpName,
							ps.s_value as SensorValue,
						   	ps.updated_date as CreatedDate
					from daily_alarms da
					inner join pump_sensor ps
						on da.serial_number = ps.serial_number
					inner join sensor_type st
					on da.type_id = st.id
					inner join pump p on ps.pump_id = p.id
					where ps.serial_number = $1
					and da.updated_date >= $2
					and da.updated_date <= $3 `

	ServiceReport = `select
						p.serial_number,
						pt.name,
						p.nick_name,
						psh.service_comment,
						psh.service_date
					from pump p
					inner join pump_service_history psh
						on p.serial_number = psh.pump_id
					inner join pump_type pt
						on p.type_id = pt.id
					where p.serial_number = $1 `
)
