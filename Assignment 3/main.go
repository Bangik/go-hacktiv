package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	go func() {
		for {
			updateJSONFile()
			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := readJSONFile()

		waterStatus := "Aman"
		if status.Water < 5 {
			waterStatus = "Aman"
		} else if status.Water >= 6 && status.Water <= 8 {
			waterStatus = "Siaga"
		} else {
			waterStatus = "Bahaya"
		}

		windStatus := "Aman"
		if status.Wind < 6 {
			windStatus = "Aman"
		} else if status.Wind >= 7 && status.Wind <= 15 {
			windStatus = "Siaga"
		} else {
			windStatus = "Bahaya"
		}

		renderTemplate(w, status, waterStatus, windStatus)
	})

	http.ListenAndServe(":8080", nil)
}

func updateJSONFile() {
	status := Status{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}

	file, err := os.Create("status.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(status)
	if err != nil {
		panic(err)
	}
}

func readJSONFile() Status {
	file, err := os.Open("status.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var status Status
	err = decoder.Decode(&status)
	if err != nil {
		panic(err)
	}

	return status
}

func renderTemplate(w http.ResponseWriter, status Status, waterStatus, windStatus string) {
	tmpl := `
    <html>
    <head>
        <title>Status</title>
        <meta http-equiv="refresh" content="15">
				<script src="https://bernii.github.io/gauge.js/dist/gauge.min.js"> </script>
				<style>
					.200x160px {
						width: 200px;
						height: 160px;
					}
					.center {
						text-align: center;
					}
				</style>
    </head>
    <body class="center">
        <h1>Status</h1>
        <p>Status Air: {{.WaterStatus}}</p>
        <p>Nilai Air: {{.Water}}</p>
				<canvas id="gauge-water" class="200x160px"></canvas>
        <p>Nilai Angin: {{.Wind}}</p>
        <p>Status Angin: {{.WindStatus}}</p>
				<canvas id="gauge-wind" class="200x160px"></canvas>
				<script>
					var opts = {
						angle: 0,
						lineWidth: 0.44,
						radiusScale: 1,
						pointer: {
							length: 0.6,
							strokeWidth: 0.035,
							color: '#000000'
						},
						limitMax: false,
						limitMin: false,
						strokeColor: '#E0E0E0',
						generateGradient: true,
						highDpiSupport: true,
						staticZones: [
							{strokeStyle: "#30B32D", min: 0, max: 8},
							{strokeStyle: "#FFDD00", min: 9, max: 15},
							{strokeStyle: "#F03E3E", min: 16, max: 100},
						],
						staticLabels: {
							font: "10px sans-serif",
							labels: [0, 8, 15, 30, 50, 70, 100],
							color: "#000000",
							fractionDigits: 0
						},
					};
					var target = document.getElementById('gauge-water');
					var gauge = new Gauge(target).setOptions(opts);
					gauge.maxValue = 100;
					gauge.setMinValue(0);
					gauge.animationSpeed = 32;
					gauge.set({{.Water}});
					var target = document.getElementById('gauge-wind');
					var gauge = new Gauge(target).setOptions(opts);
					gauge.maxValue = 100;
					gauge.setMinValue(0);
					gauge.animationSpeed = 32;
					gauge.set({{.Wind}});
				</script>
    </body>
    </html>
    `
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Water       int
		Wind        int
		WaterStatus string
		WindStatus  string
	}{
		Water:       status.Water,
		Wind:        status.Wind,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
