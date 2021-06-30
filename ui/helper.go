package ui

import "github.com/tenntenn/natureremo"

func makeRow(app *natureremo.Appliance) []string {
	var row []string

	switch app.Type {
	case natureremo.ApplianceTypeAirCon:
		if app.AirConSettings.Button == "" {
			row = []string{"ON"}
		} else {
			row = []string{"OFF"}
		}
	case natureremo.ApplianceTypeLight:
		if app.Light.State.Power == "off" {
			row = []string{"OFF"}
		} else {
			row = []string{"ON"}
		}
	case natureremo.ApplianceTypeTV:
		row = []string{string(app.TV.State.Input)}
	default:
		row = []string{"-"}
	}

	row = append(row, []string{
		app.Nickname,
		string(app.Type),
		app.Model.Name,
		app.Model.Manufacturer,
		app.Model.Country,
	}...)
	return row
}

func makeRows(apps []*natureremo.Appliance) [][]string {
	rows := make([][]string, len(apps))

	for i, app := range apps {
		row := makeRow(app)
		rows[i] = row
	}

	return rows
}
